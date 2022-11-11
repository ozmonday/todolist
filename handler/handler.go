package handler

import (
	"net/http"
	"todolists/models"
	"todolists/utility"
)

func Test(c utility.ReqRes) {
	c.Res.Write([]byte("Todolist API"))
}

func ActivityGroup(c utility.ReqRes) {
	switch c.Req.Method {
	case http.MethodGet:
		getAll(c)
		return
	case http.MethodPost:
		create(c)
		return
	default:
		c.Res.WriteHeader(http.StatusNotFound)
		return
	}
}

func getAll(c utility.ReqRes) {
	acts, err := models.SelectAllActivity(c.DB, c.Req.Context())
	if err != nil {
		c.WriteResponse(http.StatusInternalServerError, err.Error())
		return
	}

	data := []map[string]interface{}{}
	for _, v := range acts {
		data = append(data, v.Map())
	}
	c.WriteResponseJSON(http.StatusOK, "Success", data)
}

func create(c utility.ReqRes) {
	var payload utility.Payload
	var activity models.Activity

	err := c.ParseJSON(&payload)
	if err != nil {
		c.WriteResponse(http.StatusBadRequest, err.Error())
		return
	}
	activity.Parse(payload)
	activity.Insert(c.DB, c.Req.Context())
	c.WriteResponseJSON(http.StatusOK, "Success", activity.Map())
}
