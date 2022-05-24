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

func CreateEpisode(w http.ResponseWriter, r *http.Request) {

	var e models.Episode

	// Parse json body request
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		PrintErr(w, err.Error())
		return
	}

	// Add episode in the database
	episodeId, err := e.NewEpisode()
	if err != nil {
		log.Println(err)
		PrintInternalErr(w)
		return
	}

	// Send response to the client
	w.Write([]byte(fmt.Sprintf(`{"id": %d}`, episodeId)))
}

func GetEpisode(w http.ResponseWriter, r *http.Request) {

	// Get path params
	params := mux.Vars(r)

	// Convert id to int64
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		PrintErr(w, "invalid id")
		return
	}

	// Get episode from db
	episode, err := models.GetEpisodeById(id)
	if err != nil {
		log.Println("error to get episode info:", err)
		PrintInternalErr(w)
		return
	}

	// check if it's empty
	if episode.Id == 0 {
		http.Error(w, "episode not found", http.StatusNotFound)
		return
	}

	// create json response
	data, err := json.Marshal(episode)
	if err != nil {
		log.Println("error to create json response:", err)
		PrintInternalErr(w)
		return
	}

	w.Write(data)
}
