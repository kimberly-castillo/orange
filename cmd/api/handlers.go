//Filename: cmd/api/handlers.go

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/kimberly-castillo/orange/internal/data"
)

func (app *application) createSchoolHandler(w http.ResponseWriter, r *http.Request) {
	//creating a struct to hold a school that will be provided to us via the request
	//api should be able to hold post requests
	//post (write) means you want to affect the server/db
	//GET doesnt change the db
	//user will send it as a json object, so we take it, store it in go program in instance of struct in order to save to db
	var input struct {
		Name    string   `json:"name"`
		Level   string   `json:"level"`
		Contact string   `json:"contact"`
		Phone   string   `json:"phone"`
		Email   string   `json:"email"`
		Website string   `json:"website,omitempty"`
		Address string   `json:"address"`
		Mode    []string `json:"mode"`
	}
	//initialize a new json.Decorder
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		//decode returns error so check for error
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}
	//print the request
	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showSchoolHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParams(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	//fmt.Fprintf(w, "show details of school %d\n", id)
	school := data.School{
		ID:       id,
		CreateAt: time.Now(),
		Name:     "University of Belmopan",
		Level:    "University",
		Contact:  "Abel Blanco",
		Phone:    "323-4545",
		Website:  "https://uob.edu.bz",
		Address:  "17 Apple Avenue",
		Mode:     []string{"blended", "online", "face-to-face"},
		Version:  1,
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"school": school}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}
