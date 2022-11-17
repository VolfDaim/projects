package service

import (
	"projects/models"
	"projects/repository"
)

type User interface {
	Create(user models.User) (int, error)
	GetAll() ([]models.User, error)
}

type Service struct {
	User
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repos),
	}
}
