package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/zanovru/gin-rest-api/internal/config"
)

const (
	usersTable = "users"
)

func NewPostgresDB(configs *config.Configs) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		configs.Host, configs.DB.Port, configs.Username, configs.Database, configs.Password, configs.SslMode))
	if err != nil {
		return nil, err
	}

	return db, nil

}
