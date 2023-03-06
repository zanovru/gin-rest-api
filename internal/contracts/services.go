package contracts

import "github.com/zanovru/gin-rest-api/internal/models"

type AuthorizationServiceContract interface {
	RegisterUser(user models.User) (int, error)
}
