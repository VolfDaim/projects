package service

import (
	"projects/models"
	"projects/repository"
)

type TransactionService struct {
	repo repository.Transaction
}

func NewTransactionService(repo repository.Transaction) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) Send(userId int, balance int) (models.User, error) {
	return s.repo.Send(userId, balance)
}

func (s *TransactionService) Reserve(userId int, serviceId int, orderId int, balance int) error {
	return s.repo.Reserve(userId, serviceId, orderId, balance)
}

func (s *TransactionService) Confirm(userId int, serviceId int, orderId int, balance int) error {
	return s.repo.Confirm(userId, serviceId, orderId, balance)
}
