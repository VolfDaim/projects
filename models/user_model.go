package models

type User struct {
	ID      int    `json:"-" db:"id"`
	Name    string `json:"name" binding:"required"`
	Balance int    `json:"balance" binding:"required"`
}
