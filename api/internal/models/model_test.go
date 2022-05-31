package models_test

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB_DRIVER = "mysql"
var DB_USERNAME = "root"
var DB_PASSWORD = "Prova112."
var DB_HOSTNAME = "127.0.0.1"
var DB_PORT = "3306"
var DB_NAME = "StreamingScheduler"

func initDB() *sql.DB {
	var err error

	// Create db url
	dbUrl := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4",
		DB_USERNAME,
		DB_PASSWORD,
		DB_HOSTNAME,
		DB_PORT,
		DB_NAME,
	)

	log.Printf(
		"Connecting to db %s:%s/%s",
		DB_HOSTNAME,
		DB_PORT,
		DB_NAME,
	)

	// Create the database handle, confirm driver is present
	dbConn, err := sql.Open(DB_DRIVER, dbUrl)
	if err != nil {
		log.Fatal(err)
	}

	return dbConn
}
