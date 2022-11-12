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

func (a *Activity) Insert(db *sql.DB, context context.Context) error {
	var key string
	var val string
	for k, v := range a.Map() {
		if v == nil {
			continue
		}
		kind := fmt.Sprintf("%T", v)
		if len(key) == 0 && len(val) == 0 {
			key = k
			if kind == "string" {
				val = fmt.Sprintf("'%s'", v)
			} else {
				val = fmt.Sprintf("%v", v)
			}
			continue
		}

		key = fmt.Sprintf("%s, %s", key, k)
		if kind == "string" {
			val = fmt.Sprintf("%s, '%v'", val, v)
		} else {
			val = fmt.Sprintf("%s, %v", val, v)
		}
	}
	query := fmt.Sprintf("INSERT INTO activities (%s) VALUES (%s);", key, val)
	_, err := db.ExecContext(context, query)
	if err != nil {
		return err
	}
	result := `SELECT @id as "id", @created_at as "created_at", @updated_at as "updated_at";`
	row := db.QueryRowContext(context, result)
	row.Scan(&a.ID, &a.Log.Create, &a.Log.Update)
	return nil
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

func SelectActivity(db *sql.DB, context context.Context, id string) (Activity, error) {

	query := fmt.Sprintf("SELECT id, email, title, created_at, updated_at, deleted_at FROM activities WHERE id = %s;", id)
	row := db.QueryRowContext(context, query)
	r := Activity{}
	err := row.Scan(&r.ID, &r.Email, &r.Title, &r.Log.Create, &r.Log.Update, &r.Log.Delete)
	if err != nil {
		return Activity{}, err
	}
	return r, nil
}

func (a *Activity) Update(db *sql.DB, ctx context.Context, id string) error {
	var data string
	for k, v := range a.Map() {
		if v == nil {
			continue
		}

		kind := fmt.Sprintf("%T", v)
		if len(data) == 0 {
			if kind == "string" {
				data = fmt.Sprintf("%s='%v'", k, v)
			} else {
				data = fmt.Sprintf("%s=%v", k, v)
			}
			continue
		}

		if kind == "string" {
			data = fmt.Sprintf("%s, %s='%v'", data, k, v)
		} else {
			data = fmt.Sprintf("%s, %s=%v", data, k, v)
		}
	}

	query := fmt.Sprintf("UPDATE activities SET updated_at=CURRENT_TIMESTAMP, %s WHERE id=%s;", data, id)
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		return err
	}

	return nil
}

func DeleteActivity(db *sql.DB, ctx context.Context, id string) error {
	query := fmt.Sprintf("UPDATE activities SET deleted_at=CURRENT_TIMESTAMP WHERE id=%s", id)
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		return err
	}
	return nil
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
