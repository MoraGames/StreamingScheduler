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

func CreateResource(w http.ResponseWriter, r *http.Request) {

	var re models.Resource

	// Parse json body request
	err := json.NewDecoder(r.Body).Decode(&re)
	if err != nil {
		PrintErr(w, err.Error())
		return
	}

	// Add resource in the database
	resourceId, err := re.NewResource()
	if err != nil {
		log.Println(err)
		PrintInternalErr(w)
		return
	}

	// Send response to the client
	w.Write([]byte(fmt.Sprintf(`{"id": %d}`, resourceId)))
}

func GetResource(w http.ResponseWriter, r *http.Request) {

	// Get path params
	params := mux.Vars(r)

	// Convert id to int64
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		PrintErr(w, "invalid id")
		return
	}

	// Get resource from dberr
	resource, err := models.GetResourceById(id)
	if err != nil {
		log.Println("error to get resource info:", err)
		PrintInternalErr(w)
		return
	}

	// check if it's empty
	if resource.Id == 0 {
		http.Error(w, "resource not found", http.StatusNotFound)
		return
	}

	// create json response
	data, err := json.Marshal(resource)
	if err != nil {
		log.Println("error to create json response:", err)
		PrintInternalErr(w)
		return
	}

	w.Write(data)
}
