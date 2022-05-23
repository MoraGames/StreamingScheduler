package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

const ApiPrefix = "/api/v1"

func NewRouter() *mux.Router {

	// Create router
	r := mux.NewRouter()
	r.Use(enableCors) //CORS middleware

	// Define routes
	//r.HandleFunc(ApiPrefix+"/login", login).Methods(http.MethodPost)
	//r.HandleFunc(ApiPrefix+"/register", register).Methods(http.MethodPost)
	//r.HandleFunc(ApiPrefix+"/refresh", refreshToken)
	//r.HandleFunc(ApiPrefix+"/verify", verify)

	return r
}

func enableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		next.ServeHTTP(w, r)
	})
}
