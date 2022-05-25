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

func CreateEvent(w http.ResponseWriter, r *http.Request) {

	var e models.Event

	// Parse json body request
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		PrintErr(w, err.Error())
		return
	}

	// Add event in the database
	eventId, err := e.NewEvent()
	if err != nil {
		log.Println(err)
		PrintInternalErr(w)
		return
	}

	// Send response to the client
	w.Write([]byte(fmt.Sprintf(`{"id": %d}`, eventId)))
}

func GetEvent(w http.ResponseWriter, r *http.Request) {

	// Get path params
	params := mux.Vars(r)

	// Convert id to int64
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		PrintErr(w, "invalid id")
		return
	}

	// Get event from db
	event, err := models.GetEventById(id)
	if err != nil {
		log.Println("error to get event info:", err)
		PrintInternalErr(w)
		return
	}

	// check if it's empty
	if event.Id == 0 {
		http.Error(w, "event not found", http.StatusNotFound)
		return
	}

	// create json response
	data, err := json.Marshal(event)
	if err != nil {
		log.Println("error to create json response:", err)
		PrintInternalErr(w)
		return
	}

	w.Write(data)
}

func GetEvents(w http.ResponseWriter, r *http.Request) {

	// Get events from db
	events, err := models.GetEvents()
	if err != nil {
		log.Println("error to get events from db:", err)
		PrintInternalErr(w)
		return
	}

	data, err := json.Marshal(events)
	if err != nil {
		log.Println("error to create events json:", err)
		PrintInternalErr(w)
		return
	}

	w.Write(data)
}
