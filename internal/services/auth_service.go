package services

import (
	"errors"
	"github.com/zanovru/gin-rest-api/internal/contracts"
	"github.com/zanovru/gin-rest-api/internal/models"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrLoginAlreadyExists = errors.New("user with same login already exists")
)

type AuthService struct {
	repo contracts.UsersRepositoryContract
}

func NewAuthService(repo contracts.UsersRepositoryContract) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) RegisterUser(user models.User) (int, error) {
	_, err := s.repo.FindUserByLogin(user.Login)
	if err != nil {
		user.PasswordHash = hashPassword(user.PasswordHash)
		return s.repo.CreateUser(user)
	}
	return 0, ErrLoginAlreadyExists

}

func hashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes)
}
