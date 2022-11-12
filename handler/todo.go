package handler

import (
	"net/http"
	"todolists/models"
	"todolists/utility"

	"github.com/gorilla/mux"
)

func ToDo(c utility.ReqRes) {
	switch c.Req.Method {
	case http.MethodGet:
		getAllTodo(c)
		return
	case http.MethodPost:
		createTodo(c)
		return
	default:
		c.Res.WriteHeader(http.StatusNotFound)
		return
	}
}

func ToDoID(c utility.ReqRes) {
	switch c.Req.Method {
	case http.MethodGet:
		getTodoByID(c)
		return
	case http.MethodPatch:
		updateTodoByID(c)
		return
	case http.MethodDelete:
		deleteToDoByID(c)
		return
	default:
		c.Res.WriteHeader(http.StatusNotFound)
		return

	}
}

func getAllTodo(c utility.ReqRes) {
	todos, err := models.SelectAllTodo(c.DB, c.Req.Context())
	if err != nil {
		c.WriteResponseJSON(http.StatusInternalServerError, err.Error(), utility.Payload{})
		return
	}

	data := []utility.Payload{}
	for _, v := range todos {
		data = append(data, v.Map())
	}
	c.WriteResponseJSON(http.StatusOK, "Success", data)
}

func createTodo(c utility.ReqRes) {
	var payload utility.Payload
	var todo models.ToDo

	err := c.ParseJSON(&payload)
	if err != nil {
		c.WriteResponseJSON(http.StatusBadRequest, err.Error(), utility.Payload{})
		return
	}
	todo.Parse(payload)
	err = todo.Insert(c.DB, c.Req.Context())
	if err != nil {
		c.WriteResponseJSON(http.StatusInternalServerError, err.Error(), utility.Payload{})
		return
	}

	c.WriteResponseJSON(http.StatusOK, "Success", todo.Map())
}

func getTodoByID(c utility.ReqRes) {
	p := mux.Vars(c.Req)
	todo, err := models.SelectTodo(c.DB, c.Req.Context(), p["id"])
	if err != nil {
		c.WriteResponseJSON(http.StatusNotFound, err.Error(), utility.Payload{})
		return
	}
	c.WriteResponseJSON(http.StatusOK, "Success", todo.Map())
}

func updateTodoByID(c utility.ReqRes) {
	p := mux.Vars(c.Req)
	var payload utility.Payload
	var todo models.ToDo

	err := c.ParseJSON(&payload)
	if err != nil {
		c.WriteResponseJSON(http.StatusBadRequest, err.Error(), utility.Payload{})
		return
	}
	todo.Parse(payload)
	err = todo.Update(c.DB, c.Req.Context(), p["id"])
	if err != nil {
		c.WriteResponseJSON(http.StatusNotFound, err.Error(), utility.Payload{})
		return
	}
	todo, err = models.SelectTodo(c.DB, c.Req.Context(), p["id"])
	if err != nil {
		c.WriteResponseJSON(http.StatusNotFound, err.Error(), utility.Payload{})
		return
	}
	c.WriteResponseJSON(http.StatusOK, "Success", todo.Map())

}

func deleteToDoByID(c utility.ReqRes) {
	p := mux.Vars(c.Req)
	err := models.DeleteTodo(c.DB, c.Req.Context(), p["id"])
	if err != nil {
		c.WriteResponseJSON(http.StatusNotFound, err.Error(), utility.Payload{})
		return
	}
	c.WriteResponseJSON(http.StatusOK, "Success", map[string]interface{}{})
}
