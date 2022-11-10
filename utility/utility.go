package utility

import (
	"database/sql"
	"encoding/json"
	"errors"
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

func (c *ReqRes) WriteResponseJSON(statusCode int, message interface{}) error {

	data, err := json.Marshal(message)
	if err != nil {
		return err
	}
	c.Res.WriteHeader(statusCode)
	c.Res.Write(data)
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

	_, err = db.Exec(string(query))
	if err != nil {
		return err
	}

	defer file.Close()
	return nil
}
