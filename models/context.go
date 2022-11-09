package models

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DBContext struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func (db *DBContext) Connect() (*sql.DB, error) {
	data := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", db.Host, db.Port, db.User, db.Password, db.DBName)
	conn, err := sql.Open("postgres", data)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
