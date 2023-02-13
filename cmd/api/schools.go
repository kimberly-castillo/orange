//Filename: cmd/api/shcools.go

package main

import (
	"fmt"
	"net/http"
)

func (app *application) createSchoolHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Created a school...")
}

func (app *application) showSchoolHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "School displayed...")
}
