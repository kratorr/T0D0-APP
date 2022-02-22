package repository

import (
	"context"
	"fmt"
	"time"
	"todo/models"

	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"
)

type AuthPostgres struct {
	db *pgx.Conn
}

func NewAuthPostgres(db *pgx.Conn) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(u models.User) error {
	query := `INSERT INTO users (login, password) VALUES($1, $2);`
	rows, err := r.db.Query(context.Background(), query, u.Login, u.Password)
	if err != nil {
		zap.L().Sugar().Error(err.Error())

		return err
	} // TODO сраная ошибка должна быть что user already exists. !!!

	defer rows.Close()

	// fmt.Println(r.GenerateBearerToken())

	return nil
}

func (r *AuthPostgres) Authenticate(u models.User) (string, error) {
	return "", nil
}

func (r *AuthPostgres) GetUser(Login string) (models.User, error) {
	var u models.User
	// TODO а что если юзера с таким логином нет?
	query := `SELECT id, login, password FROM users WHERE login = ($1)`

	err := r.db.QueryRow(context.Background(), query, Login).Scan(&u.ID, &u.Login, &u.Password)
	if err != nil {
		zap.L().Sugar().Error(err.Error())

		return u, err
	}

	fmt.Println(u.ID)

	zap.L().Sugar().Info("ID", u.ID, "login ", u.Login, "pass hash ", u.Password)

	return u, nil
}

func (r *AuthPostgres) SaveToken(u models.User, token string) error {
	expirationDate := time.Now().AddDate(0, 0, 5)
	query := `INSERT INTO tokens (user_id, token, expiration_date) VALUES($1, $2, $3);`
	rows, err := r.db.Query(context.Background(), query, u.ID, token, expirationDate)
	if err != nil {
		zap.L().Sugar().Error(err.Error())

		return err
	}

	defer rows.Close()

	return nil
}
