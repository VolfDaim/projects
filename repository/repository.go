package repository

import (
	"github.com/jmoiron/sqlx"
	"projects/models"
)

type User interface {
	Create(user models.User) (int, error)
	GetAll() ([]models.User, error)
	GetUserBalance(userId int) (int, error)
	GetUser(userId int) (models.User, error)
	DeleteUser(userId int) error
}

type Transaction interface {
	Send(userId int, balance int) (models.User, error)
	Reserve(userId int, serviceId int, orderId int, balance int) error
	Confirm(userId int, serviceId int, orderId int, balance int) error
}

type Report interface {
	GetReport(year int, month int) (string, error)
	//DownloadReport(link string) error
}

type Repository struct {
	User
	Transaction
	Report
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User:        NewUserRepository(db),
		Transaction: NewTransactionRepository(db),
		Report:      NewReportRepository(db),
	}
}
