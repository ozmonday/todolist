package models

import (
	"database/sql"
	"fmt"
	"todolists/utility"

	_ "github.com/go-sql-driver/mysql"
)

type DBContext struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func (db *DBContext) Connect() (*sql.DB, error) {
	data := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", db.User, db.Password, db.Host, db.Port, db.DBName)
	conn, err := sql.Open("mysql", data)
	if err != nil {
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	return conn, nil
}

type Row struct {
	Create sql.NullString
	Update sql.NullString
	Delete sql.NullString
}

func (r *Row) Parse(data utility.Payload) {
	if val, ok := data["created_at"]; ok {
		r.Create.Scan(val)
	}

	if val, ok := data["updated_at"]; ok {
		r.Update.Scan(val)
	}

	if val, ok := data["deleted_at"]; ok {
		r.Delete.Scan(val)
	}
}

func (r *Row) Map() (result utility.Payload) {
	result = make(utility.Payload)
	if r.Create.Valid {
		result["created_at"] = r.Create.String
	} else {
		result["created_at"] = nil
	}

	if r.Update.Valid {
		result["updated_at"] = r.Update.String
	} else {
		result["updated_at"] = nil
	}

	if r.Delete.Valid {
		result["deleted_at"] = r.Delete.String
	} else {
		result["deleted_at"] = nil
	}

	return result
}
