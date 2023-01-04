package github

import (
	"context"
	"strconv"
	"strings"

	"github.com/cnartlu/area-service/errors"
	"github.com/cnartlu/area-service/internal/biz/city/splider/area"
	"github.com/cnartlu/area-service/internal/biz/city/splider/area/polygon"
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
			var parentID = 0
			if deep > 0 {
				var err error
				var parentArea *area.Area
				parentArea, err = g.areaUsecase.FindOneWithInstance(ctx, area.RegionIDEQ(pid), area.LevelEQ(int(deep)))
				if err != nil && !errors.IsDataNotFound(err) {
					return err
				}
				if parentArea == nil {
					// parentArea = &area.Area{
					// 	RegionID: pid,
					// 	Title:    name,
					// 	Level:    int(deep),
					// }
					// if err := g.areaUsecase.Create(ctx, parentArea); err != nil {
					// 	return err
					// }
				} else {
					parentID = parentArea.ID
				}
			}
			areaData, err := g.areaUsecase.FindOneWithInstance(ctx, area.RegionIDEQ(id), area.LevelEQ(int(deep)+1))
			if err != nil {
				if !errors.IsDataNotFound(err) {
					return err
				}
				areaData, err = g.areaUsecase.Save(ctx, &area.Area{
					ParentID:       parentID,
					RegionID:       id,
					ParentRegionID: pid,
					Title:          name,
					Level:          int(deep) + 1,
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
	areaData := []string{id, data[1], data[2], data[3]}
	if err := g.WriterByAreaData(ctx, areaData); err != nil {
		return err
	}
	geoStr := data[5]
	geoPolygonStr := data[6]
	if geoStr != "" {
		geos := strings.Split(geoStr, " ")
		area, err := g.areaUsecase.FindOneWithInstance(ctx, area.RegionIDEQ(id), area.Order("level"))
		if err != nil && !errors.IsDataNotFound(err) {
			return err
		}
		if area != nil {
			area.Lng, _ = strconv.ParseFloat(geos[0], 64)
			area.Lat, _ = strconv.ParseFloat(geos[1], 64)
			area, err = g.areaUsecase.Save(ctx, area)
			if err != nil {
				return err
			}
		}
	}
	if geoPolygonStr != "" {
		geoPolygons := strings.Split(geoPolygonStr, ",")
		for _, geoStr := range geoPolygons {
			geos := strings.Split(geoStr, " ")
			lng, _ := strconv.ParseFloat(geos[0], 64)
			lat, _ := strconv.ParseFloat(geos[1], 64)
			_, err := g.areaPolygonUsecase.Save(ctx, &polygon.Polygon{
				RegionID: id,
				Lng:      lng,
				Lat:      lat,
			})
			if err != nil {
				return err
			}
		}
	}

	return nil
}
