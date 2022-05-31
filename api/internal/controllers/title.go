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

func CreateTitle(w http.ResponseWriter, r *http.Request) {

	var t models.Title

	// Parse json body request
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		PrintErr(w, err.Error())
		return
	}

	// Add title in the database
	titleId, err := t.NewTitle()
	if err != nil {
		log.Println(err)
		PrintInternalErr(w)
		return
	}

	// Send response to the client
	w.Write([]byte(fmt.Sprintf(`{"id": %d}`, titleId)))
}

func GetTitle(w http.ResponseWriter, r *http.Request) {

	// Get path params
	params := mux.Vars(r)

	// Convert id to int64
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		PrintErr(w, "invalid id")
		return
	}

	// Get title from dberr
	title, err := models.GetTitleById(id)
	if err != nil {
		log.Println("error to get title info:", err)
		PrintInternalErr(w)
		return
	}

	// check if it's empty
	if title.Id == 0 {
		http.Error(w, "title not found", http.StatusNotFound)
		return
	}

	// create json response
	data, err := json.Marshal(title)
	if err != nil {
		log.Println("error to create json response:", err)
		PrintInternalErr(w)
		return
	}

	w.Write(data)
}

func DeleteTitle(w http.ResponseWriter, r *http.Request) {

	// Get path params
	params := mux.Vars(r)

	// Convert id to int64
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		PrintErr(w, "invalid id")
		return
	}

	err = models.DeleteTitle(id)
	if err != nil {
		PrintErr(w, err.Error())
	}

	w.Write([]byte(`{"status": "success"`))
}
