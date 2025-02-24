package service

import "databaseservice/internal/repository"



type AuthService interface {}
type Authserviceimpl struct {
	repo *repository.AuthRepositoryimpl



}
func NewAuthservice(repo *repository.AuthRepositoryimpl)*Authserviceimpl{
	return &Authserviceimpl{
		repo: repo,
	}	
}



