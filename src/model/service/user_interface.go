package service

import (
	"github.com/theus-ortiz/crud-go/src/config/rest_err"
	"github.com/theus-ortiz/crud-go/src/model"
	"github.com/theus-ortiz/crud-go/src/model/repository"
)

// Atualize a função para receber o repositório
func NewUserDomainService(repo repository.UserRepository) UserDomainService {
	return &userDomainService{
		repo: repo, // Armazene o repositório na estrutura
	}
}

type userDomainService struct {
	repo repository.UserRepository // Adicione o repositório como um campo
}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *rest_err.RestErr
	UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr
	FindUser(string) (*model.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}