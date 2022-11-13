package utility

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Payload map[string]interface{}

type ReqRes struct {
	Res    http.ResponseWriter
	Req    *http.Request
	DB     *sql.DB
	Config map[string]string
}

func (c *ReqRes) ParseJSON(payload *Payload) error {
	if c.Req.Header.Get("Content-Type") != "application/json" {
		return errors.New("request content-type is not allow")
	}

	decoder := json.NewDecoder(c.Req.Body)
	err := decoder.Decode(payload)
	if err != nil {
		return err
	}
	return nil
}

func (c *ReqRes) WriteResponseJSON(statusCode int, message string, data interface{}) error {
	var res map[string]interface{} = make(map[string]interface{})
	res["data"] = data
	res["message"] = message
	res["status"] = http.StatusText(statusCode)

	d, err := json.Marshal(res)
	if err != nil {
		return err
	}
	c.Res.Header().Add("Content-Type", "application/json")
	c.Res.WriteHeader(statusCode)
	c.Res.Write(d)
	return nil
}

func (c *ReqRes) WriteResponse(statusCode int, message string) {
	c.Res.WriteHeader(statusCode)
	c.Res.Write([]byte(message))
}

func Migration(filename string, db *sql.DB) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	query, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	fmt.Println(string(query))
	_, err = db.Exec(string(query))
	if err != nil {
		return err
	}

	fmt.Println("Minration Success")
	defer file.Close()
	return nil
}
