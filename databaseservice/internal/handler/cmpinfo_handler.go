package handler

import "databaseservice/internal/service"


type CmpinfoHandlerImpl struct {
	cmpinfo *service.CmpinfoHandlerImpl
}
type CmpinfoHandler interface {}

func NewcmpinfoHandler (service *service.CmpinfoHandlerImpl)*CmpinfoHandlerImpl{
	return &CmpinfoHandlerImpl{
		cmpinfo: service,
	}
}