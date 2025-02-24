package service

import "databaseservice/internal/repository"




type CmplainService interface {
}
type CmplainServiceImpl struct {
	repo *repository.CmplinReporImpl
}

func NewComplainserviceimpl(repo *repository.CmplinReporImpl) *CmplainServiceImpl {
	return &CmplainServiceImpl{	
		repo: repo,
	}
}