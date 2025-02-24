package handler

import "databaseservice/internal/service"

//rah ykon hna get all compalins li yjo mor ma ytna9aw ywlo rapport  ta3 complain


type CmplainHandler interface {}
type CmplainHandlerImpl struct {
	CmplainService *service.CmplainServiceImpl
}

func Newcomplainhandler( service *service.CmplainServiceImpl) *CmplainHandlerImpl {
return &CmplainHandlerImpl{
	CmplainService: service,
}
}
