package service

import (
	"projects/models"
	"projects/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(user models.User) (int, error) {
	return s.repo.Create(user)
}

func (s *UserService) GetAll() ([]models.User, error) {
	return s.repo.GetAll()
}
