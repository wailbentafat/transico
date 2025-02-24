package handler

import (
	"databaseservice/internal/service"
	"databaseservice/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CmpinfoHandlerImpl struct {
	cmpinfoService *service.CmpinfoHandlerImpl
}

type CmpinfoHandler interface {
	CreateCmpInfo(c *gin.Context)
}

func NewCmpinfoHandler(service *service.CmpinfoHandlerImpl) *CmpinfoHandlerImpl {
	return &CmpinfoHandlerImpl{
		cmpinfoService: service,
	}
}

func (h *CmpinfoHandlerImpl) CreateCmpInfo(c *gin.Context) {
	var reqBody models.CmpInfo
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cmpinfo, err := h.cmpinfoService.CreateCmpinfo(&reqBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cmpinfo)
}

