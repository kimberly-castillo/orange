//Filename: cmd/api/handlers.go

package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/kimberly-castillo/orange/internal/data"
)

func (app *application) createSchoolHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Created a school...")
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
