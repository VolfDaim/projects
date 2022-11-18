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

func (s *UserService) GetUserBalance(userId int) (int, error) {
	return s.repo.GetUserBalance(userId)
}

func (s *UserService) GetUser(userId int) (models.User, error) {
	return s.repo.GetUser(userId)
}

func (s *UserService) DeleteUser(userId int) error {
	return s.repo.DeleteUser(userId)
}
