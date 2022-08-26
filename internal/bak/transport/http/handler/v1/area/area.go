package area

import (
	"github.com/cnartlu/area-service/internal/service/area"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	areaService area.Servicer
}

func (h *Handler) Import(c *gin.Context) {
	// h.areaService.Import(c, )
}

// NewHandler 请求
func NewHandler(areaService area.Servicer) *Handler {
	return &Handler{
		areaService: areaService,
	}
}
