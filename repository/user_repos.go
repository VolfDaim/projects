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
	createUserQuery := fmt.Sprintf("INSERT INTO %s (id, balance) VALUES ($1, $2) RETURNING id", userTable)
	rows := r.db.QueryRow(createUserQuery, user.ID, user.Balance)

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

func (r *UserRepository) GetUserBalance(userId int) (int, error) {
	var user models.User
	createQuery := fmt.Sprintf("select * from %s where id=$1", userTable)
	if err := r.db.Get(&user, createQuery, userId); err != nil {
		return 0, err
	}
	return user.Balance, nil
}

func (r *UserRepository) GetUser(userId int) (models.User, error) {
	var user models.User
	createQuery := fmt.Sprintf("SELECT * from %s WHERE id = $1", userTable)
	if err := r.db.Get(&user, createQuery, userId); err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepository) DeleteUser(userId int) error {
	createQuery := fmt.Sprintf("DELETE from %s where id=$1", userTable)
	_, err := r.db.Exec(createQuery, userId)
	return err
}
