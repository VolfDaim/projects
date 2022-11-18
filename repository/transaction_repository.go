package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	models "projects/models"
)

type TransactionRepository struct {
	db *sqlx.DB
}

func NewTransactionRepository(db *sqlx.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (t TransactionRepository) Send(userId int, balance int) (models.User, error) {
	var input models.User
	if balance < 0 {
		err := errors.New("Error: trying to send negative balance")
		return input, err
	}
	createQuery := fmt.Sprintf("select * from %s where id=$1", userTable)
	if err := t.db.Get(&input, createQuery, userId); err != nil {
		return input, err
	}
	newBalance := balance + input.Balance
	createQueryUpdate := fmt.Sprintf("update %s set balance=$1 where id=$2", userTable)
	_, err := t.db.Exec(createQueryUpdate, newBalance, userId)
	if err != nil {
		return input, err
	}
	if err := t.db.Get(&input, createQuery, userId); err != nil {
		return input, err
	}
	return input, err
}

func (t TransactionRepository) Reserve(userId int, serviceId int, orderId int, balance int) error {
	var user models.User
	createQuery := fmt.Sprintf("insert into %s values($1, $2, $3, $4)", reserveTable)
	_, err := t.db.Exec(createQuery, userId, serviceId, orderId, balance)
	if err != nil {
		return err
	}
	createQueryUserSelect := fmt.Sprintf("select * from %s where id=$1", userTable)
	if err := t.db.Get(&user, createQueryUserSelect, userId); err != nil {
		return err
	}
	createQueryUserUpdate := fmt.Sprintf("update %s set balance=$1 where id=$2", userTable)
	newBalance := user.Balance - balance
	if newBalance < 0 {
		return errors.New("Unable to pay for the service: insufficient funds")
	}
	_, err = t.db.Exec(createQueryUserUpdate, newBalance, userId)
	if err != nil {
		return err
	}
	return nil
}

func (t TransactionRepository) Confirm(userId int, serviceId int, orderId int, balance int) error {
	createQuery := fmt.Sprintf("insert into %s values ($1, $2, $3, $4)", transactionTable)
	_, err := t.db.Exec(createQuery, userId, serviceId, orderId, balance)
	if err != nil {
		panic(err)
	}
	createQueryReserveDelete := fmt.Sprintf("delete from %s where id=$1 and id_service=$2 and id_order=$3 and balance=$4", reserveTable)
	if _, err := t.db.Exec(createQueryReserveDelete, userId, serviceId, orderId, balance); err != nil {
		return err
	}

	return nil
}
