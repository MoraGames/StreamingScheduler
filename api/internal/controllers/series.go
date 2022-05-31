package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/MoraGames/StreamingScheduler/core/internal/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func CreateSeries(w http.ResponseWriter, r *http.Request) {

	var s models.Series

	// Parse json body request
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		PrintErr(w, err.Error())
		return
	}

	// Add series in the database
	seriesId, err := s.NewSeries()
	if err != nil {
		log.Println(err)
		PrintInternalErr(w)
		return
	}

	// Send response to the client
	w.Write([]byte(fmt.Sprintf(`{"id": %d}`, seriesId)))
}

func GetSeries(w http.ResponseWriter, r *http.Request) {

	// Get path params
	params := mux.Vars(r)

	// Convert id to int64
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		PrintErr(w, "invalid id")
		return
	}

	// Get series from dberr
	series, err := models.GetSeriesById(id)
	if err != nil {
		log.Println("error to get series info:", err)
		PrintInternalErr(w)
		return
	}

	// check if it's empty
	if series.Id == 0 {
		http.Error(w, "series not found", http.StatusNotFound)
		return
	}

	// create json response
	data, err := json.Marshal(series)
	if err != nil {
		log.Println("error to create json response:", err)
		PrintInternalErr(w)
		return
	}

	w.Write(data)
}

func DeleteSeries(w http.ResponseWriter, r *http.Request) {

	// Get path params
	params := mux.Vars(r)

	// Convert id to int64
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		PrintErr(w, "invalid id")
		return
	}

	err = models.DeleteSeries(id)
	if err != nil {
		PrintErr(w, err.Error())
	}

	w.Write([]byte(`{"status": "success"`))
}
