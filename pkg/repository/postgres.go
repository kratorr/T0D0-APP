package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

type PostgresConfig struct {
	Host     string
	Port     int
	DBname   string
	Username string
	Password string
	SSLMode  string
}

func NewPostgresDB(config PostgresConfig) (*pgx.Conn, error) {
	params := []interface{}{config.Username, config.Password, config.Host, config.Port, config.DBname}
	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", params...)

	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	//	defer conn.Close(context.Background())

	// var name string
	// var weight int64
	// err = conn.QueryRow(context.Background(), "select name, weight from widgets where id=$1", 42).Scan(&name, &weight)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	// 	os.Exit(1)
	// }

	// fmt.Println(name, weight)
	return conn, nil
}
