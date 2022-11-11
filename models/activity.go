package models

import (
	"context"
	"database/sql"
	"fmt"
	"todolists/utility"
)

type Activity struct {
	Log   Row
	ID    sql.NullInt64
	Email sql.NullString
	Title sql.NullString
}

func (a *Activity) Insert(db *sql.DB, context context.Context) {
	var key string
	var val string
	for k, v := range a.Map() {
		if len(key) == 0 && len(val) == 0 {
			key = k
			val = fmt.Sprintf("'%s'", v)
			continue
		}

		key = fmt.Sprintf("%s, %s", key, k)
		val = fmt.Sprintf("%s, '%s'", val, v)
	}
	query := fmt.Sprintf("INSERT INTO activities (%s) VALUES (%s);", key, val)
	db.ExecContext(context, query)

	result := `SELECT @id as "id", @created_at as "created_at", @updated_at as "updated_at";`
	row := db.QueryRowContext(context, result)
	row.Scan(&a.ID, &a.Log.Create, &a.Log.Update)
}

func SelectAllActivity(db *sql.DB, context context.Context) ([]Activity, error) {
	result := []Activity{}
	query := "SELECT id, email, title, created_at, updated_at, deleted_at FROM activities;"
	rows, err := db.QueryContext(context, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		r := Activity{}
		err := rows.Scan(&r.ID, &r.Email, &r.Title, &r.Log.Create, &r.Log.Update, &r.Log.Delete)
		if err != nil {
			continue
		}
		result = append(result, r)
	}

	return result, nil
}

func (a *Activity) Parse(data utility.Payload) {
	a.Log.Parse(data)
	if val, ok := data["id"]; ok {
		a.ID.Scan(val)
	}
	if val, ok := data["email"]; ok {
		a.Email.Scan(val)
	}
	if val, ok := data["title"]; ok {
		a.Title.Scan(val)
	}
}

func (a *Activity) Map() (result utility.Payload) {
	result = make(utility.Payload)
	log := a.Log.Map()
	for k, v := range log {
		result[k] = v
	}

	if a.ID.Valid {
		result["id"] = a.ID.Int64
	}

	if a.Email.Valid {
		result["email"] = a.Email.String
	}

	if a.Title.Valid {
		result["title"] = a.Title.String
	}

	return result
}

func CreateActivity() {

}
