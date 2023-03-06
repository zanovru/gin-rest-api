package services

import (
	"github.com/zanovru/gin-rest-api/internal/contracts"
	"github.com/zanovru/gin-rest-api/internal/repositories"
)

type Services struct {
	contracts.AuthorizationServiceContract
}

func NewServices(repos *repositories.Repositories) *Services {
	return &Services{
		AuthorizationServiceContract: NewAuthService(repos.UsersRepositoryContract),
	}
}
