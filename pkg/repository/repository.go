package repository

type Auth interface {
	CreateUser() error
	Authenticate() error
}

type Repository struct {
	Auth
}

func NewRepository() *Repository {
	return &Repository{
		Auth: NewAuthPostgres(), // DB connection as param
	}
}

// type ReceiverListPostgres struct {
// 	db *pgxpool.Pool
// }

// func NewReceiverListPostgres(db *pgxpool.Pool) *ReceiverListPostgres {
// 	return &ReceiverListPostgres{db: db}
// }
