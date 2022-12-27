package github

import (
	"context"
	"strconv"
	"strings"

	"github.com/cnartlu/area-service/errors"
	"github.com/cnartlu/area-service/internal/biz/city/splider/area"
	"golang.org/x/sync/errgroup"
)

func (g *GithubUsecase) WriteByGithub(ctx context.Context, data *Github) error {
	eg, gctx := errgroup.WithContext(ctx)
	eg.SetLimit(3)
	for _, data := range data.Assets {
		data := data
		eg.Go(func() error {
			return g.LoadReleaseAsset(gctx, data, false)
		})
	}
	return eg.Wait()
}

func (g *GithubUsecase) WriterByAreaData(ctx context.Context, data []string) error {
	id := area.ToConvertRegionID(data[0])
	if id == "" {
		return errors.ErrorParamMissing("missing field `%s`", areaHeaders[0])
	}
	pid := area.ToConvertRegionID(data[1])
	if pid == "" {
		return errors.ErrorParamMissing("missing field `%s`", areaHeaders[1])
	}
	deep, err := strconv.ParseInt(data[2], 10, 8)
	if err != nil {
		return errors.ErrorParamFormat("format field `%s` error", areaHeaders[2]).WithCause(err)
	}
	name := strings.ToLower(data[3])
	{
		// 递归循环创建父级
		err := g.transaction.Transaction(ctx, func(ctx context.Context) error {
			if deep > 0 {
				var err error
				var parentArea *area.Area
				parentArea, err = g.areaUsecase.FindOneWithInstance(ctx, area.RegionIDEQ(pid), area.LevelEQ(int(deep)))
				if err != nil && !errors.IsDataNotFound(err) {
					return err
				}
				if parentArea == nil {
					parentArea = &area.Area{
						RegionID: pid,
						Title:    name,
						Level:    int(deep),
					}
					if err := g.areaUsecase.Create(ctx, parentArea); err != nil {
						return err
					}
				}
			}
			areaData, err := g.areaUsecase.FindOneWithInstance(ctx, area.RegionIDEQ(id), area.LevelEQ(int(deep)+1))
			if err != nil {
				if !errors.IsDataNotFound(err) {
					return err
				}
				areaData, err = g.areaUsecase.Save(ctx, &area.Area{
					RegionID: id,
					Title:    name,
					Level:    int(deep) + 1,
				})
				if err != nil {
					return err
				}
			}
			areaData.Title = name
			_, err = g.areaUsecase.Save(ctx, areaData)
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *GithubUsecase) WriterByGeoData(ctx context.Context, data []string) error {
	id := area.ToConvertRegionID(data[0])
	if id == "" {
		return errors.ErrorParamMissing("missing field `%s`", areaHeaders[0])
	}
	pid := area.ToConvertRegionID(data[1])
	if pid == "" {
		return errors.ErrorParamMissing("missing field `%s`", areaHeaders[1])
	}
	deep, err := strconv.ParseInt(data[2], 10, 8)
	if err != nil {
		return errors.ErrorParamFormat("format field `%s` error", areaHeaders[2]).WithCause(err)
	}
	name := strings.ToLower(data[3])
	{
		// 递归循环创建父级
		err := g.transaction.Transaction(ctx, func(ctx context.Context) error {
			if deep > 0 {
				var err error
				var parentArea *area.Area
				parentArea, err = g.areaUsecase.FindOneWithInstance(ctx, area.RegionIDEQ(pid), area.LevelEQ(int(deep)))
				if err != nil && !errors.IsDataNotFound(err) {
					return err
				}
				if parentArea == nil {
					parentArea = &area.Area{
						RegionID: pid,
						Title:    name,
						Level:    int(deep),
					}
					if err := g.areaUsecase.Create(ctx, parentArea); err != nil {
						return err
					}
				}
			}
			areaData, err := g.areaUsecase.FindOneWithInstance(ctx, area.RegionIDEQ(id), area.LevelEQ(int(deep)+1))
			if err != nil {
				if !errors.IsDataNotFound(err) {
					return err
				}
				areaData, err = g.areaUsecase.Save(ctx, &area.Area{
					RegionID: id,
					Title:    name,
					Level:    int(deep) + 1,
				})
				if err != nil {
					return err
				}
			}
			areaData.Title = name
			_, err = g.areaUsecase.Save(ctx, areaData)
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
	}
	return nil
}
