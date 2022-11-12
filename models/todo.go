package models

import (
	"context"
	"database/sql"
	"fmt"
	"todolists/utility"
)

type ToDo struct {
	Log        Row
	ID         sql.NullInt64
	Title      sql.NullString
	ActivityID sql.NullInt64
	IsActive   sql.NullBool
	Priority   sql.NullString
}

func (t *ToDo) Parse(data utility.Payload) {
	t.Log.Parse(data)
	if val, ok := data["id"]; ok {
		t.ID.Scan(val)
	}
	if val, ok := data["title"]; ok {
		t.Title.Scan(val)
	}
	if val, ok := data["is_active"]; ok {
		t.IsActive.Scan(val)
	}
	if val, ok := data["activity_group_id"]; ok {
		t.ActivityID.Scan(val)
	}
	if val, ok := data["priority"]; ok {
		t.Priority.Scan(val)
	}
}

func (t *ToDo) Map() (result utility.Payload) {
	result = make(utility.Payload)
	log := t.Log.Map()
	for k, v := range log {
		result[k] = v
	}

	if t.ID.Valid {
		result["id"] = t.ID.Int64
	}
	if t.Title.Valid {
		result["title"] = t.Title.String
	}
	if t.IsActive.Valid {
		result["is_active"] = t.IsActive.Bool
	}
	if t.ActivityID.Valid {
		result["activity_group_id"] = t.ActivityID.Int64
	}
	if t.Priority.Valid {
		result["priority"] = t.Priority.String
	}

	return result
}

func (t *ToDo) Insert(db *sql.DB, ctx context.Context) error {
	var key string
	var val string
	for k, v := range t.Map() {
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
	query := fmt.Sprintf("INSERT INTO todos (%s) VALUES (%s);", key, val)
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		return err
	}

	result := `SELECT @id_todo as "id", @is_active_todo as "is_active", @priority_todo as "priority", @created_at_todo as "created_at", @updated_at_todo as "updated_at";`
	row := db.QueryRowContext(ctx, result)
	err = row.Scan(&t.ID, &t.IsActive, &t.Priority, &t.Log.Create, &t.Log.Update)
	if err != nil {
		return err
	}

	return nil
}

func (t *ToDo) Update(db *sql.DB, ctx context.Context, id string) error {
	var data string
	for k, v := range t.Map() {
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

	query := fmt.Sprintf("UPDATE todos SET updated_at=CURRENT_TIMESTAMP, %s WHERE id=%s;", data, id)

	_, err := db.ExecContext(ctx, query)
	if err != nil {
		return err
	}

	return nil
}

func SelectAllTodo(db *sql.DB, context context.Context) ([]ToDo, error) {
	result := []ToDo{}
	query := "SELECT id, title, is_active, activity_group_id, priority, created_at, updated_at, deleted_at FROM todos;"
	rows, err := db.QueryContext(context, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		r := ToDo{}
		err := rows.Scan(&r.ID, &r.Title, &r.IsActive, &r.ActivityID, &r.Priority, &r.Log.Create, &r.Log.Update, &r.Log.Delete)
		if err != nil {
			continue
		}
		result = append(result, r)
	}

	return result, nil
}

func SelectTodo(db *sql.DB, context context.Context, id string) (ToDo, error) {
	query := fmt.Sprintf("SELECT id, title, is_active, activity_group_id, priority, created_at, updated_at, deleted_at FROM todos WHERE id = %s;", id)
	row := db.QueryRowContext(context, query)
	r := ToDo{}
	err := row.Scan(&r.ID, &r.Title, &r.IsActive, &r.ActivityID, &r.Priority, &r.Log.Create, &r.Log.Update, &r.Log.Delete)
	if err != nil {
		return ToDo{}, err
	}

	return r, nil
}

func DeleteTodo(db *sql.DB, ctx context.Context, id string) error {
	query := fmt.Sprintf("UPDATE todos SET deleted_at=CURRENT_TIMESTAMP WHERE id=%s", id)
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		return err
	}
	return nil
}
