package service

import (
	"context"

	"github.com/cnartlu/area-service/api"
	pb "github.com/cnartlu/area-service/api/v1"

	"github.com/cnartlu/area-service/internal/biz/area"
)

type AreaService struct {
	pb.UnimplementedAreaServer

	// 业务逻辑
	area *area.ManagerUsecase
}

func NewAreaService(area *area.ManagerUsecase) *AreaService {
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

func (s *AreaService) CascadeList(ctx context.Context, req *pb.CascadeListAreaRequest) (*pb.CascadeListAreaReply, error) {
	var parentData *area.Area
	var err error
	if req.GetRegionId() != "" {
		parentData, err = s.area.FindByRegionID(ctx, req.GetRegionId(), 0)
	} else if req.GetId() != 0 {
		parentData, err = s.area.FindOne(ctx, req.GetId())
	}
	if err != nil {
		return nil, err
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
	results, err := s.area.CascadeList(ctx, parentID, int(req.GetDeep()))
	if err != nil {
		return nil, err
	}
	var handlerFunc func([]*area.CascadeArea) []*pb.CascadeListAreaReply_Item
	handlerFunc = func(results []*area.CascadeArea) []*pb.CascadeListAreaReply_Item {
		items := make([]*pb.CascadeListAreaReply_Item, len(results))
		for k, result := range results {
			result := result
			item := pb.CascadeListAreaReply_Item{
				Id:       result.ID,
				RegionId: result.RegionID,
				Title:    result.Title,
				Ucfirst:  result.Ucfirst,
				Pinyin:   result.Pinyin,
				Level:    uint32(result.Level),
				Items:    make([]*pb.CascadeListAreaReply_Item, result.ChildrenNumber),
			}
			if result.ChildrenNumber > 0 {
				item.Items = handlerFunc(result.Items)
			}
			items[k] = &item
		}
		return items
	}

	xy.Items = handlerFunc(results)
	return xy, nil
}

func (s *AreaService) Create(ctx context.Context, req *pb.CreateAreaRequest) (*pb.CreateAreaReply, error) {
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
