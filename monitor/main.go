package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strconv"
	"time"
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

	// Get period event checks
	timeTicker, err := strconv.Atoi(os.Getenv("TIME_TICKER"))
	if err != nil {
		log.Fatal(err)
	}

	// Start monitoring
	log.Infoln("Starting monitoring...")
	for _ = range time.Tick(time.Second * time.Duration(timeTicker)) {
		checkDatabase()
	}
}

// checkDatabase is a function that monitoring the events start
func checkDatabase() {

	var id int64

	rows, err := dbConn.Query(`SELECT id FROM Events WHERE time <= ?`, time.Now().Unix())
	if err != nil {
		log.Error("error to monitoring events:", err.Error())
		return
	}

	//TODO: Manage multi-events result

	for rows.Next() {
		rows.Scan(&id)
	}

	// Check if there is an event to notify
	if id != 0 {
		return
	}

	// Create url
	url := os.Getenv("NOTIFY_ENDPOINT") + "?id=" + strconv.Itoa(int(id))

	// Notify request
	resp, err := http.Get(url)
	if err != nil {
		log.Error("error to call notify api:", err.Error())
		return
	}

	if resp.StatusCode != 200 {
		log.Error("error to notify event with status code:", resp.StatusCode)
	} else {
		log.Infoln("Notify event succeed")
	}
}
