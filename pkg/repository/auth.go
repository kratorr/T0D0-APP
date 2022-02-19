package repository

type AuthPostgres struct{}

func NewAuthPostgres() *AuthPostgres {
	return &AuthPostgres{}
}

func (r *AuthPostgres) CreateUser() error {
	return nil
}

func (r *AuthPostgres) Authenticate() error {
	return nil
}
