package service

import "databaseservice/internal/repository"




type Cmpinfoservice interface {}
type CmpinfoHandlerImpl struct {
	repo *repository.CmpinfoRepoImpl
}
func NewCmpinfoHandlerImpl(repo *repository.CmpinfoRepoImpl) *CmpinfoHandlerImpl {
	return &CmpinfoHandlerImpl{
		repo: repo,
	}
}