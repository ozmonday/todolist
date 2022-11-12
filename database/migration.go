package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"todolists/models"
)

func migration(filename string, db *sql.DB) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	query, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	_, err = db.Exec(string(query))
	if err != nil {
		return err
	}

	fmt.Println("Minration Success")
	defer file.Close()
	return nil
}

func main() {
	db := models.DBContext{
		Host:     os.Getenv("MYSQL_HOST"),
		Port:     os.Getenv("MYSQL_PORT"),
		User:     os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		DBName:   os.Getenv("MYSQL_DBNAME"),
	}
	conn, err := db.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	migration(os.Args[1], conn)
}