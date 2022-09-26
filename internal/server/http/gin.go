package http

import (
	"embed"
	"net/http"

	"github.com/cnartlu/area-service/internal/server/http/router"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

//go:embed page/*
var templateFS embed.FS
var (
	pageError, _ = templateFS.ReadFile("page/error.html")
)

func NewGin(grpcAddress string) (*gin.Engine, error) {
	gr, err := grpc.Dial(grpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	e := gin.Default()
	e.HandleMethodNotAllowed = true

	rootGroup := e.Group("/")

	rootGroup.Use(func(c *gin.Context) {
	})

	// 注册无路由返回
	e.NoRoute(func(c *gin.Context) {
		switch c.ContentType() {
		case "application/json":
			c.AbortWithStatusJSON(http.StatusNotFound, errors.NotFound("PageNotFound", "page not found."))
		default:
			c.Abort()
			c.Data(http.StatusNotFound, "text/html", pageError)
		}
	})

	// 注册无效的请求方法
	e.NoMethod(func(c *gin.Context) {
		switch c.ContentType() {
		case "application/json":
			c.AbortWithStatusJSON(http.StatusMethodNotAllowed, errors.NotFound("MethodNotAllowed", "request method not allowed."))
		default:
			// 405的html
			c.Abort()
			c.Data(http.StatusMethodNotAllowed, "text/html", pageError)
		}
	})
	// 注册其他路由
	router.NewArea(gr, rootGroup)
	return e, nil
}
