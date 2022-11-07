package router

import (
	"net/http"

	v1 "github.com/cnartlu/area-service/api/v1"
	"github.com/cnartlu/area-service/internal/server/http/response"
	"github.com/cnartlu/area-service/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport/http/binding"
)

type Area struct {
	s *service.AreaService
}

func (a Area) Register(g *gin.RouterGroup) {
	s := a.s
	g1 := g.Group("area")
	{
		g1.GET("/list", func(c *gin.Context) {
			param := &v1.ListAreaRequest{}
			err := binding.BindQuery(c.Request.URL.Query(), param)
			if err != nil {
				c.AbortWithStatusJSON(errors.Code(err), err)
				return
			}
			res, err := s.List(c.Request.Context(), param)
			if err != nil {
				err := errors.FromError(err)
				c.AbortWithStatusJSON(int(err.GetCode()), err)
				return
			}
			c.JSON(errors.Code(nil), response.NewSuccessDataResponse(res))
		})

		g1.GET("/cascade-list", func(c *gin.Context) {
			param := &v1.CascadeListAreaRequest{}
			err := binding.BindQuery(c.Request.URL.Query(), param)
			if err != nil {
				c.AbortWithStatusJSON(errors.Code(err), err)
				return
			}
			res, err := s.CascadeList(c.Request.Context(), param)
			if err != nil {
				err := errors.FromError(err)
				c.AbortWithStatusJSON(int(err.GetCode()), err)
				return
			}
			c.JSON(errors.Code(nil), response.NewSuccessDataResponse(res))
		})

		g1.GET("/view", func(c *gin.Context) {
			param := &v1.GetAreaRequest{}
			err := binding.BindQuery(c.Request.URL.Query(), param)
			if err != nil {
				c.AbortWithStatusJSON(errors.Code(err), err)
				return
			}
			res, err := s.View(c.Request.Context(), param)
			if err != nil {
				err := errors.FromError(err)
				c.AbortWithStatusJSON(int(err.GetCode()), err)
				return
			}
			c.JSON(errors.Code(nil), response.NewSuccessDataResponse(res))
		})

		g1.POST("/create", func(c *gin.Context) {
			res, err := s.Create(c.Request.Context(), &v1.CreateAreaRequest{})
			if err != nil {
				c.AbortWithError(http.StatusOK, err)
				return
			}
			c.JSON(http.StatusOK, response.NewSuccessDataResponse(res))
		})
		g1.POST("/update", func(c *gin.Context) {
			res, err := s.Update(c.Request.Context(), &v1.UpdateAreaRequest{})
			if err != nil {
				c.AbortWithError(http.StatusOK, err)
				return
			}
			c.JSON(http.StatusOK, response.NewSuccessDataResponse(res))
		})
		g1.POST("/delete", func(c *gin.Context) {
			res, err := s.Delete(c.Request.Context(), &v1.DeleteAreaRequest{})
			if err != nil {
				c.AbortWithError(http.StatusOK, err)
				return
			}
			c.JSON(http.StatusOK, response.NewSuccessDataResponse(res))
		})
	}
}

func NewArea(
	s *service.AreaService,
) *Area {
	return &Area{
		s: s,
	}
}
