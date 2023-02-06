package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// global variable to hold the application
// version number
const version = "1.0.0"

type config = struct {
	port int
	env  string
}

// setup dependency injection
type application struct {
	config config //variable name "config" of type config
	logger *log.Logger
}

func main() {
	var cfg config
	//get the argumeents for the user for the server configuration
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment(development|staging|production)")
	flag.Parse()
	//create a logger
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	//create an object if type application
	app := &application{
		config: cfg,
		logger: logger, //for terminal

	}
	//create a route/mutliplexer
	//mutliplexer is a data structure that works like a map
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler) //the route and the function

	//create our server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.port),
		Handler: mux,

		IdleTimeout: time.Minute,      //inactive connective //or end connection after 1 min
		ReadTimeout: 10 * time.Second, //time to read request body or header
		//ddos attack can occur if it doesnt stop reading
		WriteTimeout: 10 * time.Second,
	}
	//start our server
	logger.Printf("Starting %s Server on port %d", cfg.env, cfg.port)
	err := srv.ListenAndServe()
	logger.Fatal(err) //if error then print out the error
}
