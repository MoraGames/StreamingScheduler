package controllers

import (
	"io"
	"log"
	"net/http"
)

func GetCurrentUser(w http.ResponseWriter, r *http.Request) {

	// create request
	req, err := http.NewRequest(http.MethodGet, "http://auth:5000/api/v1/info", nil)
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
