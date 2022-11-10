package handler

import (
	"fmt"
	"net/http"
	"todolists/models"
	"todolists/utility"
)

func Test(c utility.ReqRes) {
	fmt.Println(c.Req.URL)
}

func ActivityGroup(c utility.ReqRes) {
	switch c.Req.Method {
	case http.MethodGet:
		fmt.Println("GET")
		getAll(c)
		return
	case http.MethodPost:
		fmt.Println("POST")
		create(c)
		return
	default:
		c.Res.WriteHeader(http.StatusNotFound)
		return
	}
}

func getAll(c utility.ReqRes) {

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

	go fmt.Println(activity.Map())
}
