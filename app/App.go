package app

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"todolists/models"
	"todolists/utility"

	"github.com/gorilla/mux"
)

type Engine struct {
	route  *mux.Router
	db     *sql.DB
	config map[string]string
}

func NewEngine(db models.DBContext) Engine {
	conn, _ := db.Connect()
	utility.Migration(os.Getenv("QUERY"), conn)
	return Engine{
		route:  mux.NewRouter(),
		db:     conn,
		config: map[string]string{},
	}
}

func (a Engine) Run(port string) {
	server := http.Server{
		Handler: a.route,
		Addr:    fmt.Sprintf(":%s", port),
	}
	server.ListenAndServe()
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

// default

var defaultApp Engine

func init() {

}
