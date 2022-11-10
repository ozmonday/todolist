package models

import (
	"database/sql"
	"todolists/utility"
)

type Activity struct {
	Log   Row
	ID    sql.NullString
	Email sql.NullString
	Title sql.NullString
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
		result["id"] = a.ID.String
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
