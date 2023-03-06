package contracts

import "github.com/zanovru/gin-rest-api/internal/models"

type UsersRepositoryContract interface {
	CreateUser(user models.User) (int, error)
	FindUserByLogin(login string) (models.User, error)
}
