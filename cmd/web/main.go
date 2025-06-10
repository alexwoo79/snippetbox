package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"

	"github.com/alexwoo79/snippetbox/internal/models"
	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	logger   *slog.Logger
	snippets *models.SnippetModel
}

func main() {
	addr := flag.String("addr", ":4000", "Http network address")
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL data source name")
	flag.Parse()
	// logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	// To keep the main() function tidy I've put the code for creating a connection // pool into the separate openDB() function below. We pass openDB() the DSN // from the command-line flag.

	db, err := openDB(*dsn)
	if err != nil {

		logger.Error(err.Error())

		os.Exit(1)
	}

	// We also defer a call to db.Close(), so that the connection pool is closed // before the main() function exits.

	defer db.Close()
	app := &application{
		logger:   logger,
		snippets: &models.SnippetModel{DB: db},
	}

	logger.Info("Starting server", "addr", *addr)
	// using the http.ListenAndServe() function to start a new web server
	// port and mux
	err = http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}

// The openDB() function wraps sql.Open() and returns a sql.DB connection pool // for a given DSN.

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
