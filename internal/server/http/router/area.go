package router

import (
	"net/http"

	v1 "github.com/cnartlu/area-service/api/v1"
	"github.com/cnartlu/area-service/internal/server/http/response"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport/http/binding"
	"google.golang.org/grpc"
)

func NewArea(
	grpc *grpc.ClientConn,
	g *gin.RouterGroup,
) {
	client := v1.NewAreaClient(grpc)
	g1 := g.Group("area")
	{
		g1.GET("/list", func(c *gin.Context) {
			param := &v1.ListAreaRequest{}
			err := binding.BindQuery(c.Request.URL.Query(), param)
			if err != nil {
				c.AbortWithStatusJSON(errors.Code(err), err)
				return
			}
			res, err := client.List(c.Request.Context(), param)
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
			res, err := client.CascadeList(c.Request.Context(), param)
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
			res, err := client.View(c.Request.Context(), param)
			if err != nil {
				err := errors.FromError(err)
				c.AbortWithStatusJSON(int(err.GetCode()), err)
				return
			}
			c.JSON(errors.Code(nil), response.NewSuccessDataResponse(res))
		})

		g1.POST("/create", func(c *gin.Context) {
			res, err := client.Create(c.Request.Context(), &v1.CreateAreaRequest{})
			if err != nil {
				c.AbortWithError(http.StatusOK, err)
				return
			}
			c.JSON(http.StatusOK, response.NewSuccessDataResponse(res))
		})
		g1.POST("/update", func(c *gin.Context) {
			res, err := client.Update(c.Request.Context(), &v1.UpdateAreaRequest{})
			if err != nil {
				c.AbortWithError(http.StatusOK, err)
				return
			}
			c.JSON(http.StatusOK, response.NewSuccessDataResponse(res))
		})
		g1.POST("/delete", func(c *gin.Context) {
			res, err := client.Delete(c.Request.Context(), &v1.DeleteAreaRequest{})
			if err != nil {
				c.AbortWithError(http.StatusOK, err)
				return
			}
			c.JSON(http.StatusOK, response.NewSuccessDataResponse(res))
		})
	}
}
