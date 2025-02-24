package service

import (
	"databaseservice/internal/repository"
	Types "databaseservice/pkj/types"
)

type AuthService interface{}
type Authserviceimpl struct {
	repo *repository.AuthRepositoryimpl
}

func (a *Authserviceimpl) Register(user *Types.User) any {
	panic("unimplemented")
}

func (a *Authserviceimpl) Login(user *Types.User) any {
	panic("unimplemented")
}
func NewAuthservice(repo *repository.AuthRepositoryimpl) *Authserviceimpl {
	return &Authserviceimpl{
		repo: repo,
	}
}
