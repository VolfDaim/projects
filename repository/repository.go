package repository

import (
	"github.com/jmoiron/sqlx"
	"projects/models"
)

type User interface {
	Create(user models.User) (int, error)
	GetAll() ([]models.User, error)
}

type Repository struct {
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User: NewUserRepository(db),
	}
}
