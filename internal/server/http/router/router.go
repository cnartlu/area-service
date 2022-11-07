package router

import "github.com/gin-gonic/gin"

type Router interface {
	Register(g *gin.RouterGroup)
}

func NewRouter(
	a *Area,
) []Router {
	var routers = []Router{}
	routers = append(
		routers,
		a,
	)
	return routers
}
