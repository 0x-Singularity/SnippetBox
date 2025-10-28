package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// Define an application struct to hold the application wide dependencies (like the structured logger)
type application struct {
	logger *slog.Logger
}

func main() {
	//Define a new command line flag, defaults to 4000 if not specified
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL data source name")
	flag.Parse()

	// Introducing structured logging, writes to stdout and nil as second parameter keeps default settings
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer db.Close()

	//Initialize new instance of the application struct that contains the logger
	app := &application{
		logger: logger,
	}

	mux := http.NewServeMux()

	//This file server allows us to serve files out of the ui/static directory.
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	//Other appliation routes
	//swapped route declarations to use the applications structs methods as handler functions
	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)
	mux.HandleFunc("GET /snippet/create", app.snippetCreate)
	mux.HandleFunc("POST /snippet/create", app.snippetCreatePost)

	//use the info() method to log the start server message at "info" severity
	logger.Info("starting server", slog.String("addr", ":4000"))

	err = http.ListenAndServe(*addr, mux)
	// Use Error() method to log any error message returned by http.ListenAndServe() at Error severity
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
