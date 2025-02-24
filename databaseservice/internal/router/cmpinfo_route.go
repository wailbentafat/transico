package router

import (
	
	"github.com/gin-gonic/gin"
)

type Cmpinforoute struct{}

func (r *Cmpinforoute) RegisterRoutes(router *gin.Engine, handler *handler.CmpinfoHandlerImpl) {
	cmpinfo := router.Group("/cmpinfo")
	cmpinfo.POST("/", handler.CreateCmpInfo)
	cmpinfo.GET("/", handler.GetCmpInfo)
	cmpinfo.GET("/:id", handler.GetCmpInfoById)
	cmpinfo.PUT("/:id", handler.UpdateCmpInfo)
	cmpinfo.DELETE("/:id", handler.DeleteCmpInfo)
	
}
