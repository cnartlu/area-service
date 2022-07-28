package db

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strings"

	"github.com/cnartlu/area-service/internal/component/ent"
	"github.com/cnartlu/area-service/pkg/component/log"
	"go.uber.org/zap"
)

type txFunc func(tx *ent.Tx) error

type DB struct {
	logger   *log.Logger
	client   *ent.Client
	clients  map[string]*ent.Client
	cleanups map[string]func()
	keys     []string
}

func (db *DB) DB() *ent.Client {
	return db.client
}

func (db *DB) Ent(name string) *ent.Client {
	return db.clients[name]
}

func (db *DB) EntTx(name string, ctx context.Context, fn txFunc) error {
	tx, err := db.Ent(name).Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("rolling back transaction: %w", rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}

func (db *DB) DBTx(ctx context.Context, fn txFunc) error {
	tx, err := db.DB().Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("rolling back transaction: %w", rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}

func New(logger *log.Logger, config *Config) (*DB, func(), error) {
	if config == nil {
		return nil, nil, errors.New("db config can't be empty, please check your configuration")
	}
	obj := DB{
		logger:   logger,
		client:   nil,
		clients:  make(map[string]*ent.Client),
		cleanups: make(map[string]func()),
		keys:     []string{},
	}
	if config != nil && config.Connections != nil {
		for key, value := range config.Connections {
			key = strings.ToLower(strings.TrimSpace(key))
			client, cleanup, err := NewEnt(value, logger)
			if err != nil {
				obj.logger.Error("db connection error", zap.String("db", key), zap.Error(err))
				continue
			}
			obj.cleanups[key] = cleanup
			obj.clients[key] = client
			obj.keys = append(obj.keys, key)
		}
		if len(obj.keys) < 1 {
			return nil, nil, errors.New("db connections fail, please check your configuration")
		}
		// 默认的数据库
		defaultdb := config.Default
		if defaultdb != "" {
			defaultdb = strings.ToLower(strings.TrimSpace(defaultdb))
		}
		if client, ok := obj.clients[defaultdb]; ok {
			obj.client = client
		}
		if obj.client == nil && defaultdb != "db" {
			if client, ok := obj.clients["db"]; ok {
				obj.client = client
			}
		}
		// 随机取值
		if obj.client == nil {
			key := rand.Intn(len(obj.keys))
			obj.client = obj.clients[obj.keys[key]]
		}
	}
	cleanup := func() {
		for _, cleanup := range obj.cleanups {
			cleanup()
		}
	}
	return &obj, cleanup, nil
}
