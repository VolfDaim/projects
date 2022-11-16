package service

import "projects/repository"

type User interface {
}

type Service struct {
	User
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
