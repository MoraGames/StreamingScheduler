package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

// database connection variable
var dbConn *sql.DB

// logger variable
var log = logrus.New()

func init() {
	var err error

	// Setting logger
	log.Out = os.Stdout

	log.Infoln(os.Getenv("DB_DRIVER"))

	// Create db url
	dbUrl := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOSTNAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	log.Infof(
		"Connecting to db %s:%s/%s",
		os.Getenv("DB_DRIVER"),
		":",
		os.Getenv("DB_HOSTNAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	// Create the database handle, confirm driver is present
	dbConn, err = sql.Open(os.Getenv("DB_DRIVER"), dbUrl)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	port := os.Getenv("PORT")

	srv := &http.Server{
		Handler: NewRouter(),
		Addr:    ":" + port,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Infoln("Starting server on :" + port)

	log.Fatal(srv.ListenAndServe())
}
