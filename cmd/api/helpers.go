// Filename: cmd/api/helpers.go

package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) readIDParams(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("Invalid ID parameter")
	}
	return id, nil
}

func (app *application) WriteJSON(w http.ResponseWriter, status int, data interface{}, headers http.Header) error {
	//convert data to json format
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}
	js = append(js, '\n')
	//any headers that were sent
	for key, value := range headers {
		w.Header()[key] = value
	}
	//add header information
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(js))
	return nil
}
