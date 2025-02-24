package handler

import "databaseservice/internal/service"

type CmpDataHandler interface{}
type CmpHandler struct {
	cmpservice *service.CmpdataServiceImpl
}
func Newcmpdatahandler(cmpservice *service.CmpdataServiceImpl) *CmpHandler {
	return &CmpHandler{
		cmpservice: cmpservice,
	}	
}
func ()
