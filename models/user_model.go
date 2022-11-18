package models

type User struct {
	ID      int `json:"id" db:"id"`
	Balance int `json:"balance" db:"balance"`
}
