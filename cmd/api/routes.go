//Filename: cmd/api/routes.go

package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//this function returns the router
func (app *application) routes() *httprouter.Router {
	//create a new router
	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/schools", app.createSchoolHandler)
	router.HandlerFunc(http.MethodGet, "/v1/schools/:id", app.showSchoolHandler)
	//return the router
	return router
}
