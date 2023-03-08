//Filename: cmd/api/healthcheck.go
package main

import (
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	//js := `{"status":"available", "environment":%q, "version": %q}`
	//js = fmt.Sprintf(js, app.config.env, version)

	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	err := app.WriteJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
		return
	}
	// fmt.Fprintf(w, "status: available\n")
	// fmt.Fprintf(w, "environment: %s \n ", app.config.env)
	// fmt.Fprintf(w, "version: %s\n", version)
}
