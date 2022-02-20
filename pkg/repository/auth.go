package repository

import (
	"context"
	"fmt"
	"math/rand"
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

	fmt.Println(r.GenerateBearerToken())

	return nil
}

func (r *AuthPostgres) Authenticate(u models.User) (string, error) {
	return "", nil
}

func (r *AuthPostgres) GetUser(ID string) (models.User, error) {
	var u models.User

	query := `SELECT id, login, password FROM users WHERE login = ($1)`

	err := r.db.QueryRow(context.Background(), query, ID).Scan(&u.ID, &u.Login, &u.Password)
	if err != nil {
		zap.L().Sugar().Error(err.Error())

		return u, err
	}
	zap.L().Sugar().Info("login ", u.Login, "pass hash ", u.Password)

	return u, nil
}

func (r *AuthPostgres) CreateToken(u models.User) (string, error) {
	return r.GenerateBearerToken(), nil
}

func (r *AuthPostgres) GenerateBearerToken() string {
	rand.Seed(time.Now().UnixNano())

	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, 40)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}
