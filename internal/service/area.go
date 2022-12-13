package service

import (
	"context"

	"github.com/cnartlu/area-service/api"
	pb "github.com/cnartlu/area-service/api/v1"
	"golang.org/x/sync/errgroup"

	"github.com/cnartlu/area-service/internal/biz/area"
)

type AreaService struct {
	pb.UnimplementedAreaServer

	// 业务逻辑
	area *area.AreaUsecase
}

func NewAreaService(area *area.AreaUsecase) *AreaService {
	return &AreaService{
		area: area,
	}
}

func (s *AreaService) List(ctx context.Context, req *pb.ListAreaRequest) (*pb.ListAreaReply, error) {
	results, err := s.area.List(ctx, area.FindListParam{
		ParentID: uint64(req.GetParentId()),
		RegionID: req.GetRegionId(),
		Level:    int(req.GetLevel()),
		Keyword:  req.GetKw(),
		Order:    req.GetOrder(),
	})
	if err != nil {
		return nil, err
	}
	reply := &pb.ListAreaReply{}
	for _, result := range results {
		result := result
		reply.Items = append(reply.Items, &pb.ListAreaReply_Item{
			Id:         result.ID,
			RegionId:   result.RegionID,
			Title:      result.Title,
			Ucfirst:    result.Ucfirst,
			Pinyin:     result.Pinyin,
			CityCode:   result.CityCode,
			ZipCode:    result.ZipCode,
			Level:      uint32(result.Level),
			UpdateTime: uint64(result.UpddateAt.Unix()),
		})
	}
	return reply, nil
}

func (s *AreaService) View(ctx context.Context, req *pb.GetAreaRequest) (*pb.GetAreaReply, error) {
	var (
		id     = req.GetId()
		result *area.Area
		err    error
	)
	if id > 0 {
		result, err = s.area.FindOne(ctx, id)
	} else {
		result, err = s.area.FindByRegionID(ctx, req.GetRegionId(), int(req.GetLevel()))
	}
	if err != nil {
		return nil, err
	}
	return &pb.GetAreaReply{
		Id:         result.ID,
		RegionId:   result.RegionID,
		Title:      result.Title,
		Ucfirst:    result.Ucfirst,
		Pinyin:     result.Pinyin,
		CityCode:   result.CityCode,
		ZipCode:    result.ZipCode,
		Level:      uint32(result.Level),
		CreateTime: uint64(result.CreateAt.Unix()),
		UpdateTime: uint64(result.UpddateAt.Unix()),
		Parent:     nil,
	}, nil
}

// CascadeList 级联列表
func (s *AreaService) CascadeList(ctx context.Context, req *pb.CascadeListAreaRequest) (*pb.CascadeListAreaReply, error) {
	var parentData *area.Area

	{
		var err error
		if req.GetRegionId() != "" {
			parentData, err = s.area.FindByRegionID(ctx, req.GetRegionId(), int(req.GetLevel()))
		} else if req.GetId() != 0 {
			parentData, err = s.area.FindOne(ctx, req.GetId())
		}
		if err != nil {
			return nil, err
		}
	}

	var parentID uint64
	xy := &pb.CascadeListAreaReply{}
	if parentData != nil {
		parentID = parentData.ID
		xy.Parent = &pb.CascadeListAreaReply_Item{
			Id:       parentData.ID,
			RegionId: parentData.RegionID,
			Title:    parentData.Title,
			Ucfirst:  parentData.Ucfirst,
			Pinyin:   parentData.Pinyin,
			Level:    uint32(parentData.Level),
		}
	}

	{
		eg, cancelCtx := errgroup.WithContext(ctx)
		var (
			handlerFunc func(parentID uint64, deep int) ([]*pb.CascadeListAreaReply_Item, error)
			maxDeep     = int(req.GetDeep())
		)
		if maxDeep < area.AREA_MIN_LEVEL {
			maxDeep = area.AREA_MAX_LEVEL
		}
		handlerFunc = func(parentID uint64, deep int) ([]*pb.CascadeListAreaReply_Item, error) {
			results, err := s.area.List(cancelCtx, area.FindListParam{ParentID: parentID})
			if err != nil {
				return nil, err
			}
			items := make([]*pb.CascadeListAreaReply_Item, len(results))
			for idx := 0; idx < len(results); idx++ {
				idx := idx
				result := results[idx]
				item := &pb.CascadeListAreaReply_Item{
					Id:       result.ID,
					RegionId: result.RegionID,
					Title:    result.Title,
					Lat:      float32(result.Lat),
					Lng:      float32(result.Lng),
					Ucfirst:  result.Ucfirst,
					Pinyin:   result.Pinyin,
					Level:    uint32(result.Level),
					Items:    make([]*pb.CascadeListAreaReply_Item, 0),
				}
				items[idx] = item
				if deep < maxDeep && result.Level < area.AREA_MAX_LEVEL {
					eg.Go(func() error {
						items, err := handlerFunc(result.ID, deep+1)
						if err != nil {
							return err
						}
						item.Items = items
						return nil
					})
				}
			}
			return items, nil
		}
		results, err := handlerFunc(parentID, 1)
		if err != nil {
			return nil, err
		}
		if err := eg.Wait(); err != nil {
			return nil, err
		}
		xy.Items = results
	}

	return xy, nil
}

func (s *AreaService) Create(ctx context.Context, req *pb.CreateAreaRequest) (*pb.CreateAreaReply, error) {
	s.area.Create(ctx, area.CreateParam{})
	return &pb.CreateAreaReply{}, nil
}

func (s *AreaService) Update(ctx context.Context, req *pb.UpdateAreaRequest) (*pb.UpdateAreaReply, error) {
	return &pb.UpdateAreaReply{}, nil
}

func (s *AreaService) Delete(ctx context.Context, req *pb.DeleteAreaRequest) (*pb.DeleteAreaReply, error) {
	var ids = req.GetIds()
	if req.GetId() > 0 {
		ids = append(ids, req.GetId())
	}
	for _, id := range ids {
		id := id
		if id < 1 {
			continue
		}
	}
	if len(ids) < 1 {
		return nil, api.ErrorParamMissing("identify parameter missing")
	}
	err := s.area.Delete(ctx, ids...)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteAreaReply{}, nil
}
