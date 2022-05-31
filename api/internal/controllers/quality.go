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

func CreateQuality(w http.ResponseWriter, r *http.Request) {

	var q models.Quality

	// Parse json body request
	err := json.NewDecoder(r.Body).Decode(&q)
	if err != nil {
		PrintErr(w, err.Error())
		return
	}

	// Add quality in the database
	qualityid, err := q.NewQuality()
	if err != nil {
		log.Println(err)
		PrintInternalErr(w)
		return
	}

	// Send response to the client
	w.Write([]byte(fmt.Sprintf(`{"id": %d}`, qualityid)))
}

func GetQuality(w http.ResponseWriter, r *http.Request) {

	// Get path params
	params := mux.Vars(r)

	// Convert id to int64
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		PrintErr(w, "invalid id")
		return
	}

	// Get quality from dberr
	quality, err := models.GetQualityById(id)
	if err != nil {
		log.Println("error to get quality info:", err)
		PrintInternalErr(w)
		return
	}

	// check if it's empty
	if quality.Id == 0 {
		http.Error(w, "quality not found", http.StatusNotFound)
		return
	}

	// create json response
	data, err := json.Marshal(quality)
	if err != nil {
		log.Println("error to create json response:", err)
		PrintInternalErr(w)
		return
	}

	w.Write(data)
}

func DeleteQuality(w http.ResponseWriter, r *http.Request) {

	// Get path params
	params := mux.Vars(r)

	// Convert id to int64
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		PrintErr(w, "invalid id")
		return
	}

	err = models.DeleteQuality(id)
	if err != nil {
		PrintErr(w, err.Error())
	}

	w.Write([]byte(`{"status": "success"`))
}
