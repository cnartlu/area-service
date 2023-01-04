package data

import (
	"context"
	"fmt"
	"time"

	"ariga.io/entcache"
	"github.com/cnartlu/area-service/component/database"
	biztransaction "github.com/cnartlu/area-service/internal/biz/transaction"
	"github.com/cnartlu/area-service/internal/data/ent"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v4/stdlib"
	"go.uber.org/zap"
)

var _ biztransaction.Transaction = (*Data)(nil)

type txCtxKey struct{}

type Data struct {
	logger *zap.Logger
	ent    *ent.Client
	rds    *redis.Client
}

// newTxContext returns a new context with the given Tx attached.
func (d *Data) newTxContext(ctx context.Context, tx *ent.Tx) context.Context {
	return context.WithValue(ctx, txCtxKey{}, tx)
}

// txFromContext returns a Tx stored inside a context, or nil if there isn't one.
func (d *Data) txFromContext(ctx context.Context) *ent.Tx {
	tx, _ := ctx.Value(txCtxKey{}).(*ent.Tx)
	return tx
}

func (d *Data) GetRedis() *redis.Client {
	return d.rds
}

func (d *Data) GetClient(ctx context.Context) *ent.Client {
	tx := d.txFromContext(ctx)
	if tx == nil {
		return d.ent
	}
	return tx.Client()
}

func (d *Data) WithCacheContext(ctx context.Context, ttl *int) context.Context {
	if ttl != nil {
		t := *ttl
		if t < 0 {
			return entcache.Skip(ctx)
		} else if t > 0 {
			return entcache.WithTTL(ctx, time.Duration(t)*time.Second)
		} else {
			return entcache.WithTTL(ctx, 0)
		}
	}
	return ctx
}

func (d *Data) Transaction(ctx context.Context, fn func(ctx context.Context) error) error {
	tx := d.txFromContext(ctx)
	if tx != nil {
		return fn(ctx)
	}
	tx, err := d.ent.Tx(ctx)
	if err != nil {
		return fmt.Errorf("starting transaction: %w", err)
	}
	if err = fn(d.newTxContext(ctx, tx)); err != nil {
		if err2 := tx.Rollback(); err2 != nil {
			return fmt.Errorf("rolling back transaction: %v (original error: %w)", err2, err)
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}

func NewData(
	logger *zap.Logger,
	rds *redis.Client,
	config *database.Config,
) (*Data, func(), error) {
	result := Data{
		logger: logger.WithOptions(zap.AddCallerSkip(1)),
		rds:    rds,
	}
	client, cleanup, err := result.LoadEntDatabase(config)
	if err != nil {
		return nil, nil, err
	}
	result.ent = client
	return &result, cleanup, err
}
