//Filename: cmd/api/healthcheck.go
package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	js := `{"status":"available", "environment":%q, "version": %q}`
	js = fmt.Sprintf(js, app.config.env, version)
	//add header information
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(js))
	// fmt.Fprintf(w, "status: available\n")
	// fmt.Fprintf(w, "environment: %s \n ", app.config.env)
	// fmt.Fprintf(w, "version: %s\n", version)
}
