package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

const ApiPrefix = "/api/v1"

func NewRouter() *mux.Router {

	// Create router
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc(ApiPrefix+"/login", login).Methods(http.MethodPost)
	r.HandleFunc(ApiPrefix+"/register", register).Methods(http.MethodPost)
	//r.HandleFunc(ApiPrefix + "/refresh", refreshToken)
	//r.HandleFunc(ApiPrefix + "/verify", verify)

	return r
}
