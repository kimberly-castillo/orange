package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// global variable to hold the application
// version number
const version = "1.0.0"

type config = struct {
	port int
	env  string
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
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
	flag.StringVar(&cfg.db.dsn, "dsn", os.Getenv("LEMON_DB_DSN"), "PostgreSQL DSN")
	flag.IntVar(&cfg.db.maxOpenConns, "max-open-conns", 25, "PostgreSQL Max Connections")
	flag.IntVar(&cfg.db.maxIdleConns, "max-idle-conns", 25, "PostgreSQL Idle Time")
	flag.StringVar(&cfg.db.maxIdleTime, "max-idle-time", "15m", "PostgreSQL max connection Idle Time")
	flag.Parse()
	//create a logger
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	//setup the database connection pool
	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}
	//close database connection pool
	defer db.Close()

	//create an object if type application
	app := &application{
		config: cfg,
		logger: logger, //for terminal

	}

	//create our server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.port),
		Handler: app.routes(),

		IdleTimeout: time.Minute,      //inactive connective //or end connection after 1 min
		ReadTimeout: 10 * time.Second, //time to read request body or header
		//ddos attack can occur if it doesnt stop reading
		WriteTimeout: 10 * time.Second,
	}
	//start our server
	logger.Printf("Starting %s Server on port %d", cfg.env, cfg.port)
	err = srv.ListenAndServe()
	logger.Fatal(err) //if error then print out the error
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("pgx", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.db.maxOpenConns)
	db.SetMaxIdleConns(cfg.db.maxIdleConns)
	duration, err := time.ParseDuration(cfg.db.maxIdleTime)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxIdleTime(duration)

	//create a context
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	//ping the database
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}
