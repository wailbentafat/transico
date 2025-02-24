package service

import "databaseservice/internal/repository"



type CmpdataService interface {}
type CmpdataServiceImpl struct {
	repo *repository.CmpdataRepositoryImpl
}
func NewCmpdataServiceImpl (repo *repository.CmpdataRepositoryImpl)*CmpdataServiceImpl{
	return &CmpdataServiceImpl{
		repo: repo,
	}
}