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

func CreateLanguage(w http.ResponseWriter, r *http.Request) {

	var l models.Language

	// Parse json body request
	err := json.NewDecoder(r.Body).Decode(&l)
	if err != nil {
		PrintErr(w, err.Error())
		return
	}

	// Add language in the database
	languageId, err := l.NewLanguage()
	if err != nil {
		log.Println(err)
		PrintInternalErr(w)
		return
	}

	// Send response to the client
	w.Write([]byte(fmt.Sprintf(`{"id": %d}`, languageId)))
}

func GetLanguage(w http.ResponseWriter, r *http.Request) {

	// Get path params
	params := mux.Vars(r)

	// Convert id to int64
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		PrintErr(w, "invalid id")
		return
	}

	// Get language from dberr
	language, err := models.GetLanguageById(id)
	if err != nil {
		log.Println("error to get language info:", err)
		PrintInternalErr(w)
		return
	}

	// check if it's empty
	if language.Id == 0 {
		http.Error(w, "language not found", http.StatusNotFound)
		return
	}

	// create json response
	data, err := json.Marshal(language)
	if err != nil {
		log.Println("error to create json response:", err)
		PrintInternalErr(w)
		return
	}

	w.Write(data)
}
