package main

import (
	"todolists/app"
	"todolists/handler"
	"todolists/models"
)

var engine app.Engine

func initRoute() {
	engine.HandleFunc("/activity-groups", handler.Test)
	engine.HandleFunc("/activity-groups/:id", handler.Test)
	engine.HandleFunc("/todo-items", handler.Test)
	engine.HandleFunc("/todo-items/:id", handler.Test)

}

func init() {
	db := models.DBContext{
		Host:     "18.140.60.34",
		Port:     "5432",
		User:     "mita",
		Password: "mita2022",
		DBName:   "waste",
	}
	engine = app.NewEngine(db)
	initRoute()
}

func main() {
	//database setting
	engine.AddConfig("images_path", "/")
	engine.Run("3030")
}
