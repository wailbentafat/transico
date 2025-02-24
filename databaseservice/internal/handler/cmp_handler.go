package handler

import (
	"databaseservice/internal/models"
	"databaseservice/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CmpDataHandler interface{}
type CmpHandler struct {
	cmpservice *service.CmpdataServiceImpl
}
func Newcmpdatahandler(cmpservice *service.CmpdataServiceImpl) *CmpHandler {
	return &CmpHandler{
		cmpservice: cmpservice,
	}	
}
func (h *CmpHandler) CreateCmpData(c *gin.Context) {
	var cmp models.CmpData
	if err := c.ShouldBindJSON(&cmp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.cmpservice.CreateCmpData(&cmp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "heyy"})
		return
	}
	c.JSON(http.StatusCreated, cmp)
}
