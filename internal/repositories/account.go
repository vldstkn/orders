package repositories

import (
	"orders/internal/domain"
	"orders/pkg/db"
)

type AccountRepository struct {
	DB *db.DB
}

func NewAccountRepository(db *db.DB) *AccountRepository {
	return &AccountRepository{
		DB: db,
	}
}

func (repo *AccountRepository) Create(user *domain.User) (int, error) {
	var id int
	err := repo.DB.QueryRow("INSERT INTO users (email, password, name) VALUES ($1, $2, $3) RETURNING id",
		user.Email, user.Password, user.Name).Scan(&id)

	if err != nil {
		return -1, err
	}
	return id, nil
}
func (repo *AccountRepository) FindByEmail(email string) *domain.User {
	var user domain.User
	err := repo.DB.Get(&user, "SELECT * FROM users WHERE email=$1", email)
	if err != nil {
		return nil
	}
	return &user
}
