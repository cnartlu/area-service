package service

import (
	"context"

	pb "github.com/cnartlu/area-service/api/v1"
	"github.com/go-kratos/kratos/v2/errors"

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
		result, err = s.area.ViewWithIDEQ(ctx, id)
	} else {
		regionId := req.GetRegionId()
		if regionId == "" {
			return nil, errors.BadRequest("PARAMS_EMPTY", "参数不能为空")
		}
		result, err = s.area.ViewWithRegionID(ctx, regionId, int(req.GetLevel()))
	}
	if err != nil {
		return nil, errors.NotFound("NOT_FOUND", err.Error())
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
	return &pb.CascadeListAreaReply{}, nil
}

func (s *AreaService) Create(ctx context.Context, req *pb.CreateAreaRequest) (*pb.CreateAreaReply, error) {
	return &pb.CreateAreaReply{}, nil
}

func (s *AreaService) Update(ctx context.Context, req *pb.UpdateAreaRequest) (*pb.UpdateAreaReply, error) {
	return &pb.UpdateAreaReply{}, nil
}

func (s *AreaService) Delete(ctx context.Context, req *pb.DeleteAreaRequest) (*pb.DeleteAreaReply, error) {
	return &pb.DeleteAreaReply{}, nil
}
