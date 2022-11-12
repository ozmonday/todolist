package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"todolists/app"
	"todolists/handler"
	"todolists/models"
	"todolists/utility"
)

var engine app.Engine

func initRoute() {
	engine.HandleFunc("/activity-groups", handler.ActivityGroup)
	engine.HandleFunc("/activity-groups/{id}", handler.ActivityGroupID)
	engine.HandleFunc("/todo-items", handler.ToDo)
	engine.HandleFunc("/todo-items/{id}", handler.ToDoID)
}

func initMigrate(db models.DBContext) error {
	conn, err := db.Connect()
	if err != nil {
		return err
	}
	utility.Migration("/usr/src/app/database/create_table_activities.sql", conn)
	utility.Migration("/usr/src/app/database/create_table_todos.sql", conn)
	utility.Migration("/usr/src/app/database/trigger_insert_activities.sql", conn)
	utility.Migration("/usr/src/app/database/trigger_insert_todos.sql", conn)
	defer conn.Close()
	return nil
}

func main() {
	runtime.GOMAXPROCS(3)

	db := models.DBContext{
		Host:     os.Getenv("MYSQL_HOST"),
		Port:     os.Getenv("MYSQL_PORT"),
		User:     os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		DBName:   os.Getenv("MYSQL_DBNAME"),
	}
	fmt.Println(db)
	for {
		err := initMigrate(db)
		if err != nil {
			log.Println(err)
		} else {
			break
		}
	}
	engine = app.NewEngine(db)
	initRoute()
	fmt.Println("Server runing on port:", os.Getenv("PORT"))
	if err := engine.Run(os.Getenv("PORT")); err != nil {
		log.Println(err)
	}
}
