package main

import (
	"os"
	"runtime"
	"todolists/app"
	"todolists/handler"
	"todolists/models"
)

var engine app.Engine

func initRoute() {
	engine.HandleFunc("/activity-groups", handler.ActivityGroup)
	engine.HandleFunc("/activity-groups/{id}", handler.ActivityGroup)
	engine.HandleFunc("/todo-items", handler.Test)
	engine.HandleFunc("/todo-items/{id}", handler.Test)
}

func init() {
	db := models.DBContext{
		Host:     os.Getenv("MYSQL_HOST"),
		Port:     os.Getenv("MYSQL_PORT"),
		User:     os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		DBName:   os.Getenv("MYSQL_DBNAME"),
	}
	engine = app.NewEngine(db)
	initRoute()
}

func main() {

	runtime.GOMAXPROCS(3)
	engine.Run(os.Getenv("PORT"))
}
