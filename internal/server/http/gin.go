package http

import (
	"github.com/cnartlu/area-service/internal/server/http/router"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewGin(grpcAddress string) (*gin.Engine, error) {
	e := gin.New()
	rootGroup := e.Group("/")
	gr, err := grpc.Dial(grpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	router.NewArea(gr, rootGroup)
	return e, nil
}
