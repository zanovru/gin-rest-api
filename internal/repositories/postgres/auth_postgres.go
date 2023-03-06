package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/zanovru/gin-rest-api/internal/models"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user models.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (login, password_hash, name) values ($1, $2, $3) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.Login, user.PasswordHash, user.Name)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) FindUserByLogin(login string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE login=$1", usersTable)
	err := r.db.Get(&user, query, login)

	return user, err
}
