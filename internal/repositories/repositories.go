package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/zanovru/gin-rest-api/internal/contracts"
	"github.com/zanovru/gin-rest-api/internal/repositories/postgres"
)

type Repositories struct {
	contracts.UsersRepositoryContract
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		UsersRepositoryContract: postgres.NewAuthPostgres(db)}
}
