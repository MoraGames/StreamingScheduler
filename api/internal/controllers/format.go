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

func CreateFormat(w http.ResponseWriter, r *http.Request) {

	var f models.Format

	// Parse json body request
	err := json.NewDecoder(r.Body).Decode(&f)
	if err != nil {
		PrintErr(w, err.Error())
		return
	}

	// Add format in the database
	formatId, err := f.NewFormat()
	if err != nil {
		log.Println(err)
		PrintInternalErr(w)
		return
	}

	// Send response to the client
	w.Write([]byte(fmt.Sprintf(`{"id": %d}`, formatId)))
}

func GetFormat(w http.ResponseWriter, r *http.Request) {

	// Get path params
	params := mux.Vars(r)

	// Convert id to int64
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		PrintErr(w, "invalid id")
		return
	}

	// Get format from dberr
	format, err := models.GetFormatById(id)
	if err != nil {
		log.Println("error to get format info:", err)
		PrintInternalErr(w)
		return
	}

	// check if it's empty
	if format.Id == 0 {
		http.Error(w, "format not found", http.StatusNotFound)
		return
	}

	// create json response
	data, err := json.Marshal(format)
	if err != nil {
		log.Println("error to create json response:", err)
		PrintInternalErr(w)
		return
	}

	w.Write(data)
}
