// Code generated by protoc-gen-go-gin. DO NOT EDIT.
// versions:
// - protoc-gen-go-gin v0.0.1
// - protoc             v3.19.4
// source: api/manage/v1/area.proto

package v1

import (
	context "context"
	errors "errors"
	gin "github.com/gin-gonic/gin"
	binding "github.com/gin-gonic/gin/binding"
	protojson "google.golang.org/protobuf/encoding/protojson"
	ioutil "io/ioutil"
	http "net/http"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
// var _ = new(context.Context)
// const _ = binding.MIMEHTML
// const _ = http.StatusOK
// const _ = gin.ContextKey
// const _ = ioutil.Discard
// const _ = protojson.MarshalOptions{ DiscardUnknown: true }
// var _ = errors.New("area unknow error")

const AreaResponseDataKey = "_gin-gonic/gin/responseDataKey"

type AreaGinServer interface {
	CascadeListArea(context.Context, *CascadeListAreaRequest) (*CascadeListAreaReply, error)
	CreateArea(context.Context, *CreateAreaRequest) (*CreateAreaReply, error)
	DeleteArea(context.Context, *DeleteAreaRequest) (*DeleteAreaReply, error)
	GetArea(context.Context, *GetAreaRequest) (*GetAreaReply, error)
	ListArea(context.Context, *ListAreaRequest) (*ListAreaReply, error)
	UpdateArea(context.Context, *UpdateAreaRequest) (*UpdateAreaReply, error)
}

func RegisterAreaGinServer(s *gin.RouterGroup, srv AreaGinServer) {
	s.POST("/area", _Area_CreateArea0_Gin_Handler(srv))
	s.PUT("/area/:id", _Area_UpdateArea0_Gin_Handler(srv))
	s.DELETE("/area/:id", _Area_DeleteArea0_Gin_Handler(srv))
	s.GET("/area/:id", _Area_GetArea0_Gin_Handler(srv))
	s.GET("/area/list", _Area_ListArea0_Gin_Handler(srv))
	s.GET("/area/list/cascade", _Area_CascadeListArea0_Gin_Handler(srv))
}

func _Area_CreateArea0_Gin_Handler(srv AreaGinServer) gin.HandlerFunc {
	return func(c *gin.Context) {
		var in CreateAreaRequest
		var err error
		b := binding.Default(c.Request.Method, c.ContentType())
		switch b {
		case binding.Form:
			if err := c.Request.ParseForm(); err != nil {
				c.Error(err)
				return
			}
			if _, err := c.MultipartForm(); err != nil && !errors.Is(err, http.ErrNotMultipart) {
				c.Error(err)
				return
			}
			err = binding.MapFormWithTag(&in, c.Request.Form, "json")
		case binding.JSON:
			var body []byte
			if cb, ok := c.Get(gin.BodyBytesKey); ok {
				if cbb, ok := cb.([]byte); ok {
					body = cbb
				}
			}
			if body == nil {
				body, err = ioutil.ReadAll(c.Request.Body)
				if err != nil {
					c.Error(err)
					return
				}
				c.Set(gin.BodyBytesKey, body)
			}
			err = protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(body, &in)
		default:
			err = c.MustBindWith(&in, b)
		}
		if err != nil {
			c.Error(err)
			return
		}
		out, err := srv.CreateArea(c.Request.Context(), &in)
		if err != nil {
			c.Error(err)
			return
		}
		c.Set(AreaResponseDataKey, out)
	}
}

func _Area_UpdateArea0_Gin_Handler(srv AreaGinServer) gin.HandlerFunc {
	return func(c *gin.Context) {
		var in UpdateAreaRequest
		if err := c.BindUri(&in); err != nil {
			c.Error(err)
			return
		}
		var err error
		b := binding.Default(c.Request.Method, c.ContentType())
		switch b {
		case binding.Form:
			if err := c.Request.ParseForm(); err != nil {
				c.Error(err)
				return
			}
			if _, err := c.MultipartForm(); err != nil && !errors.Is(err, http.ErrNotMultipart) {
				c.Error(err)
				return
			}
			err = binding.MapFormWithTag(&in, c.Request.Form, "json")
		case binding.JSON:
			var body []byte
			if cb, ok := c.Get(gin.BodyBytesKey); ok {
				if cbb, ok := cb.([]byte); ok {
					body = cbb
				}
			}
			if body == nil {
				body, err = ioutil.ReadAll(c.Request.Body)
				if err != nil {
					c.Error(err)
					return
				}
				c.Set(gin.BodyBytesKey, body)
			}
			err = protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(body, &in)
		default:
			err = c.MustBindWith(&in, b)
		}
		if err != nil {
			c.Error(err)
			return
		}
		out, err := srv.UpdateArea(c.Request.Context(), &in)
		if err != nil {
			c.Error(err)
			return
		}
		c.Set(AreaResponseDataKey, out)
	}
}

func _Area_DeleteArea0_Gin_Handler(srv AreaGinServer) gin.HandlerFunc {
	return func(c *gin.Context) {
		var in DeleteAreaRequest
		if err := c.BindUri(&in); err != nil {
			c.Error(err)
			return
		}
		values := c.Request.URL.Query()
		if err := binding.MapFormWithTag(&in, values, "json"); err != nil {
			c.Error(err)
			return
		}
		out, err := srv.DeleteArea(c.Request.Context(), &in)
		if err != nil {
			c.Error(err)
			return
		}
		c.Set(AreaResponseDataKey, out)
	}
}

func _Area_GetArea0_Gin_Handler(srv AreaGinServer) gin.HandlerFunc {
	return func(c *gin.Context) {
		var in GetAreaRequest
		if err := c.BindUri(&in); err != nil {
			c.Error(err)
			return
		}
		values := c.Request.URL.Query()
		if err := binding.MapFormWithTag(&in, values, "json"); err != nil {
			c.Error(err)
			return
		}
		out, err := srv.GetArea(c.Request.Context(), &in)
		if err != nil {
			c.Error(err)
			return
		}
		c.Set(AreaResponseDataKey, out)
	}
}

func _Area_ListArea0_Gin_Handler(srv AreaGinServer) gin.HandlerFunc {
	return func(c *gin.Context) {
		var in ListAreaRequest
		values := c.Request.URL.Query()
		if err := binding.MapFormWithTag(&in, values, "json"); err != nil {
			c.Error(err)
			return
		}
		out, err := srv.ListArea(c.Request.Context(), &in)
		if err != nil {
			c.Error(err)
			return
		}
		c.Set(AreaResponseDataKey, out)
	}
}

func _Area_CascadeListArea0_Gin_Handler(srv AreaGinServer) gin.HandlerFunc {
	return func(c *gin.Context) {
		var in CascadeListAreaRequest
		values := c.Request.URL.Query()
		if err := binding.MapFormWithTag(&in, values, "json"); err != nil {
			c.Error(err)
			return
		}
		out, err := srv.CascadeListArea(c.Request.Context(), &in)
		if err != nil {
			c.Error(err)
			return
		}
		c.Set(AreaResponseDataKey, out)
	}
}
