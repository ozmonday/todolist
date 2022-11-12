package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"todolists/models"
	"todolists/utility"

	"github.com/NYTimes/gziphandler"
	"github.com/gorilla/mux"
)

type Engine struct {
	route  *mux.Router
	db     *sql.DB
	config map[string]string
}

func NewEngine(db models.DBContext) Engine {
	log.Println("Connect to database...")
	conn, err := db.Connect()
	if err != nil {
		log.Println(err.Error())
	}

	return Engine{
		route:  mux.NewRouter(),
		db:     conn,
		config: map[string]string{},
	}
}

func (a *Engine) Run(port string) error {

	defer a.db.Close()
	server := http.Server{
		Handler: gziphandler.GzipHandler(a.route),
		Addr:    fmt.Sprintf(":%s", port),
	}
	if err := server.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

func (a *Engine) AddConfig(key string, value string) {
	a.config[key] = value
}

func (a *Engine) HandleFunc(path string, handler func(c utility.ReqRes)) {
	a.route.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		handler(utility.ReqRes{
			Res:    w,
			Req:    r,
			DB:     a.db,
			Config: a.config,
		})

	})
}
