package main

import (
	"github.com/gorilla/mux"
)

const ApiPrefix = "/api/v1"

func NewRouter() *mux.Router {

	// Create router
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc(ApiPrefix + "/login", login)
	r.HandleFunc(ApiPrefix + "/register", register)
	//r.HandleFunc(ApiPrefix + "/refresh", refreshToken)
	//r.HandleFunc(ApiPrefix + "/verify", verify)

	return r
}
