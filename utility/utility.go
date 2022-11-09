package utility

import (
	"database/sql"
	"net/http"
)

type ReqRes struct {
	Res    http.ResponseWriter
	Req    *http.Request
	DB     *sql.DB
	Config map[string]string
}
