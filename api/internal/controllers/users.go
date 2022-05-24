package controllers

import (
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"strconv"
)

func GetCurrentUser(w http.ResponseWriter, r *http.Request) {

	// Get path params
	params := mux.Vars(r)

	// Convert id to int64
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		PrintErr(w, "invalid id")
		return
	}

	// create request
	req, err := http.NewRequest(http.MethodGet, "http://auth.streamtv.it/api/v1/info", nil)
	if err != nil {
		log.Println("error to create the request", err)
		PrintInternalErr(w)
		return
	}
	req.Header.Set("Authorization", r.Header.Get("Authorization"))

	// do the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("error to do the request", err)
		PrintInternalErr(w)
		return
	}
	defer resp.Body.Close()

	// Check response status code
	if resp.StatusCode != http.StatusOK {
		log.Println("request error with status code:", resp.StatusCode)
		http.Error(w, "request error", resp.StatusCode)
		return
	}

	// Read the data
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("error to read the auth response:", err.Error())
		PrintInternalErr(w)
		return
	}

	w.Write(content)
}
