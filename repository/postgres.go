package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	userTable = "users"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func ConnectDB(con Config) (*sqlx.DB, error) {
	config := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		con.Host, con.Port, con.Username, con.Password, con.DBName, con.SSLMode)
	db, error := sqlx.Open("postgres", config)
	if error != nil {
		return nil, error
	}
	return db, nil
}
