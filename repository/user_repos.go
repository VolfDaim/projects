package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"projects/models"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user models.User) (int, error) {
	var id int
	createUserQuery := fmt.Sprintf("INSERT INTO %s (name, balance) VALUES ($1, $2) RETURNING id", userTable)
	rows := r.db.QueryRow(createUserQuery, user.Name, user.Balance)

	if err := rows.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UserRepository) GetAll() ([]models.User, error) {
	var lists []models.User
	createQuery := "SELECT * from users"
	err := r.db.Select(&lists, createQuery)

	return lists, err
}
