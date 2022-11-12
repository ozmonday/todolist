package handler

import (
	"net/http"
	"todolists/models"
	"todolists/utility"

	"github.com/gorilla/mux"
)

func Test(c utility.ReqRes) {
	c.Res.Write([]byte("Todolist API"))
}

func ActivityGroup(c utility.ReqRes) {
	switch c.Req.Method {
	case http.MethodGet:
		getAllActivity(c)
		return
	case http.MethodPost:
		createActivity(c)
		return
	default:
		c.Res.WriteHeader(http.StatusNotFound)
		return
	}
}

func ActivityGroupID(c utility.ReqRes) {
	switch c.Req.Method {
	case http.MethodGet:
		getActivityByID(c)
		return
	case http.MethodPatch:
		updateActivityByID(c)
		return
	case http.MethodDelete:
		deleteActivityByID(c)
		return
	default:
		c.Res.WriteHeader(http.StatusNotFound)
		return
	}
}

func getActivityByID(c utility.ReqRes) {
	p := mux.Vars(c.Req)
	act, err := models.SelectActivity(c.DB, c.Req.Context(), p["id"])
	if err != nil {
		c.WriteResponseJSON(http.StatusNotFound, err.Error(), utility.Payload{})
		return
	}
	c.WriteResponseJSON(http.StatusOK, "Success", act.Map())
}

func updateActivityByID(c utility.ReqRes) {
	p := mux.Vars(c.Req)
	var payload utility.Payload
	var act models.Activity

	err := c.ParseJSON(&payload)
	if err != nil {
		c.WriteResponseJSON(http.StatusBadRequest, err.Error(), utility.Payload{})
		return
	}
	act.Parse(payload)
	err = act.Update(c.DB, c.Req.Context(), p["id"])
	if err != nil {
		c.WriteResponseJSON(http.StatusNotFound, err.Error(), utility.Payload{})
		return
	}
	act, err = models.SelectActivity(c.DB, c.Req.Context(), p["id"])
	if err != nil {
		c.WriteResponseJSON(http.StatusNotFound, err.Error(), utility.Payload{})
		return
	}
	c.WriteResponseJSON(http.StatusOK, "Success", act.Map())
}

func deleteActivityByID(c utility.ReqRes) {
	p := mux.Vars(c.Req)
	err := models.DeleteActivity(c.DB, c.Req.Context(), p["id"])
	if err != nil {
		c.WriteResponseJSON(http.StatusNotFound, err.Error(), utility.Payload{})
		return
	}
	c.WriteResponseJSON(http.StatusOK, "Success", utility.Payload{})
}

func getAllActivity(c utility.ReqRes) {
	acts, err := models.SelectAllActivity(c.DB, c.Req.Context())
	if err != nil {
		c.WriteResponseJSON(http.StatusInternalServerError, err.Error(), utility.Payload{})
		return
	}

	data := []utility.Payload{}
	for _, v := range acts {
		data = append(data, v.Map())
	}
	c.WriteResponseJSON(http.StatusOK, "Success", data)
}

func createActivity(c utility.ReqRes) {
	var payload utility.Payload
	var activity models.Activity

	err := c.ParseJSON(&payload)
	if err != nil {
		c.WriteResponseJSON(http.StatusBadRequest, err.Error(), utility.Payload{})
		return
	}
	activity.Parse(payload)
	err = activity.Insert(c.DB, c.Req.Context())
	if err != nil {
		c.WriteResponseJSON(http.StatusInternalServerError, err.Error(), utility.Payload{})
		return
	}
	c.WriteResponseJSON(http.StatusOK, "Success", activity.Map())
}
