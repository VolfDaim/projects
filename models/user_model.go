package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	ID      uint   `gorm:"primary_key"`
	Name    string `gorm:"column:name"`
	Balance int    `gorm:"column:balance"`
}
