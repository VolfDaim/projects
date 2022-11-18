package service

import (
	"projects/models"
	"projects/repository"
)

type User interface {
	Create(user models.User) (int, error)
	GetAll() ([]models.User, error)
	GetUserBalance(userId int) (int, error)
	GetUser(idUser int) (models.User, error)
	DeleteUser(userId int) error
}

type Transaction interface {
	Send(userId int, amount int) (models.User, error)
	Reserve(userId int, serviceId int, orderId int, amount int) error
	Confirm(userId int, serviceId int, orderId int, amount int) error
}

type Report interface {
	GetReport(year int, month int) (string, error)
}

type Service struct {
	User
	Transaction
	Report
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User:        NewUserService(repos),
		Transaction: NewTransactionService(repos),
		Report:      NewReportService(repos),
	}
}
