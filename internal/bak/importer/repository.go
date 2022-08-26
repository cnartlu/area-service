package importer

import (
	"context"
	"strconv"
	"strings"

	"github.com/cnartlu/area-service/internal/component/ent"
	"github.com/cnartlu/area-service/internal/component/ent/area"
	pstrings "github.com/cnartlu/area-service/pkg/utils/strings"
	"github.com/mozillazg/go-pinyin"
)

type Row struct {
	ID, Pid, Deep, Name, PinyinPrefix, Pinyin, ExtID, ExtName string
}

type Dao interface {
	ImportWithRow(ctx context.Context, row Row) error
}

type Repository struct {
	ent *ent.Client
}

func (r *Repository) getOrCreate(ctx context.Context, row Row) (*ent.Area, error) {
	deep, _ := strconv.ParseInt(row.Deep, 10, 64)
	level := deep + 1
	model, err := r.ent.Area.Query().
		Where(area.RegionIDEQ(row.ID), area.LevelEQ(uint8(level))).
		First(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return nil, err
		}
		var pinyinStr = strings.Join(pinyin.LazyPinyin(row.Name, pinyin.NewArgs()), " ")
		var ucfirst = ""
		if pinyinStr != "" {
			ucfirst = strings.ToUpper(pinyinStr)[:1]
		}
		model, err = r.ent.Area.Create().
			SetRegionID(row.ID).
			SetLevel(uint8(level)).
			SetTitle(row.Name).
			SetPinyin(pinyinStr).
			SetUcfirst(ucfirst).
			Save(ctx)
	}
	return model, err
}

func (r *Repository) ImportWithRow(ctx context.Context, row Row) error {
	id, err := strconv.ParseInt(strings.TrimRight(row.ID, "0 "), 10, 64)
	if err != nil {
		return err
	}
	pid, err := strconv.ParseInt(strings.TrimLeft(row.Pid, "0 "), 10, 64)
	if err != nil {
		return err
	}
	idStr := pstrings.Pad(id, 10, "0", pstrings.STR_PAD_RIGHT)
	deep, _ := strconv.ParseInt(row.Deep, 10, 64)
	level := deep + 1
	var pinyinStr = strings.Join(pinyin.LazyPinyin(row.Name, pinyin.NewArgs()), " ")
	var ucfirst = ""
	if pinyinStr != "" {
		ucfirst = strings.ToUpper(pinyinStr)[:1]
	}
	var (
		model       *ent.Area
		parentModel *ent.Area
	)
	if deep > 0 {
		pr := Row{
			ID:   pstrings.Pad(pid, 10, "0", pstrings.STR_PAD_RIGHT),
			Name: row.Name,
			Deep: strconv.FormatInt(level-2, 10),
		}
		parentModel, err = r.getOrCreate(ctx, pr)
		if !ent.IsNotFound(err) {
			return err
		}
	}
	model, err = r.ent.Area.Query().
		Where(area.RegionIDEQ(idStr), area.LevelEQ(uint8(level))).
		First(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return err
		}
		model, err = r.ent.Area.Create().
			SetRegionID(idStr).
			SetLevel(uint8(level)).
			SetPinyin(pinyinStr).
			SetUcfirst(ucfirst).
			Save(ctx)
	} else {
		model, err = model.Update().
			SetTitle(row.Name).
			SetPinyin(pinyinStr).
			SetUcfirst(ucfirst).
			Save(ctx)
	}
	if err != nil {
		return err
	}
	if parentModel != nil {
		_, err = model.Update().
			SetParentID(parentModel.ID).
			Save(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func NewRepository(ent *ent.Client) *Repository {
	return &Repository{
		ent: ent,
	}
}
