package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"todolists/app"
	"todolists/handler"
	"todolists/models"
)

var engine app.Engine

func initRoute() {
	engine.HandleFunc("/activity-groups", handler.ActivityGroup)
	engine.HandleFunc("/activity-groups/{id}", handler.ActivityGroupID)
	engine.HandleFunc("/todo-items", handler.ToDo)
	engine.HandleFunc("/todo-items/{id}", handler.ToDoID)
}
func main() {
	runtime.GOMAXPROCS(3)

	// db := models.DBContext{
	// 	Host:     os.Getenv("MYSQL_HOST"),
	// 	Port:     os.Getenv("MYSQL_PORT"),
	// 	User:     os.Getenv("MYSQL_USER"),
	// 	Password: os.Getenv("MYSQL_PASSWORD"),
	// 	DBName:   os.Getenv("MYSQL_DBNAME"),
	// }

	db := models.DBContext{
		Host:     "localhost",
		Port:     "3306",
		User:     "root",
		Password: "cakradara",
		DBName:   "todolist",
	}
	engine = app.NewEngine(db)
	initRoute()
	fmt.Println("Server runing on port:", os.Getenv("PORT"))
	if err := engine.Run(os.Getenv("PORT")); err != nil {
		log.Println(err)
	}
}
