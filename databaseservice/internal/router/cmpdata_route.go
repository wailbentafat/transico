package router

import (
	"github.com/gin-gonic/gin"
	"databaseservice/internal/handler"
)

type Cmpdataroute struct {
	handler *handler.CmpHandler
}

func NewCmpdataroute(handler *handler.CmpHandler) *Cmpdataroute {
	return &Cmpdataroute{handler: handler}
}

func (r *Cmpdataroute) RegisterRoutes(router *gin.Engine) {
	cmpdata := router.Group("/cmpdata")
	{
		cmpdata.POST("/", r.handler.CreateCmpData)
	
	}
}

