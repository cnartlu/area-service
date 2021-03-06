// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.2.1

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type AreaHTTPServer interface {
	CascadeListArea(context.Context, *CascadeListAreaRequest) (*CascadeListAreaReply, error)
	GetArea(context.Context, *GetAreaRequest) (*GetAreaReply, error)
	ListArea(context.Context, *ListAreaRequest) (*ListAreaReply, error)
}

func RegisterAreaHTTPServer(s *http.Server, srv AreaHTTPServer) {
	r := s.Route("/")
	r.GET("/area/list/{parent_level}/{parent_region_id}", _Area_ListArea0_HTTP_Handler(srv))
	r.GET("/area/list", _Area_ListArea1_HTTP_Handler(srv))
	r.GET("/area/{level}/{region_id}", _Area_GetArea0_HTTP_Handler(srv))
	r.GET("/area/cascade-list", _Area_CascadeListArea0_HTTP_Handler(srv))
}

func _Area_ListArea0_HTTP_Handler(srv AreaHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListAreaRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.v1.Area/ListArea")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListArea(ctx, req.(*ListAreaRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListAreaReply)
		return ctx.Result(200, reply)
	}
}

func _Area_ListArea1_HTTP_Handler(srv AreaHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListAreaRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.v1.Area/ListArea")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListArea(ctx, req.(*ListAreaRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListAreaReply)
		return ctx.Result(200, reply)
	}
}

func _Area_GetArea0_HTTP_Handler(srv AreaHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetAreaRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.v1.Area/GetArea")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetArea(ctx, req.(*GetAreaRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetAreaReply)
		return ctx.Result(200, reply)
	}
}

func _Area_CascadeListArea0_HTTP_Handler(srv AreaHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CascadeListAreaRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.v1.Area/CascadeListArea")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CascadeListArea(ctx, req.(*CascadeListAreaRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CascadeListAreaReply)
		return ctx.Result(200, reply)
	}
}

type AreaHTTPClient interface {
	CascadeListArea(ctx context.Context, req *CascadeListAreaRequest, opts ...http.CallOption) (rsp *CascadeListAreaReply, err error)
	GetArea(ctx context.Context, req *GetAreaRequest, opts ...http.CallOption) (rsp *GetAreaReply, err error)
	ListArea(ctx context.Context, req *ListAreaRequest, opts ...http.CallOption) (rsp *ListAreaReply, err error)
}

type AreaHTTPClientImpl struct {
	cc *http.Client
}

func NewAreaHTTPClient(client *http.Client) AreaHTTPClient {
	return &AreaHTTPClientImpl{client}
}

func (c *AreaHTTPClientImpl) CascadeListArea(ctx context.Context, in *CascadeListAreaRequest, opts ...http.CallOption) (*CascadeListAreaReply, error) {
	var out CascadeListAreaReply
	pattern := "/area/cascade-list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.v1.Area/CascadeListArea"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AreaHTTPClientImpl) GetArea(ctx context.Context, in *GetAreaRequest, opts ...http.CallOption) (*GetAreaReply, error) {
	var out GetAreaReply
	pattern := "/area/{level}/{region_id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.v1.Area/GetArea"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AreaHTTPClientImpl) ListArea(ctx context.Context, in *ListAreaRequest, opts ...http.CallOption) (*ListAreaReply, error) {
	var out ListAreaReply
	pattern := "/area/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.v1.Area/ListArea"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
