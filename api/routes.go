package main

import (
	"github.com/MoraGames/StreamingScheduler/core/internal/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

const ApiPrefix = "/v1"

func NewRouter() *mux.Router {

	// Create router
	r := mux.NewRouter()
	r.Use(enableCors) //CORS middleware

	// Episode routes
	r.HandleFunc(ApiPrefix+"/episodes", controllers.CreateEpisode).Methods(http.MethodPost)
	r.HandleFunc(ApiPrefix+"/episodes/{id}", controllers.GetEpisode).Methods(http.MethodGet)

	// Event routes
	r.HandleFunc(ApiPrefix+"/events", controllers.CreateEvent).Methods(http.MethodPost)
	r.HandleFunc(ApiPrefix+"/events/{id}", controllers.GetEvent).Methods(http.MethodGet)

	// Format routes
	r.HandleFunc(ApiPrefix+"/formats", controllers.CreateFormat).Methods(http.MethodPost)
	r.HandleFunc(ApiPrefix+"/formats/{id}", controllers.GetFormat).Methods(http.MethodGet)

	// Language routes
	r.HandleFunc(ApiPrefix+"/languages", controllers.CreateLanguage).Methods(http.MethodPost)
	r.HandleFunc(ApiPrefix+"/languages/{id}", controllers.GetLanguage).Methods(http.MethodGet)

	// Quality routes
	r.HandleFunc(ApiPrefix+"/qualities", controllers.CreateQuality).Methods(http.MethodPost)
	r.HandleFunc(ApiPrefix+"/qualities/{id}", controllers.GetQuality).Methods(http.MethodGet)

	// Resource routes
	r.HandleFunc(ApiPrefix+"/resources", controllers.CreateResource).Methods(http.MethodPost)
	r.HandleFunc(ApiPrefix+"/resources/{id}", controllers.GetResource).Methods(http.MethodGet)

	// Series routes
	r.HandleFunc(ApiPrefix+"/series", controllers.CreateSeries).Methods(http.MethodPost)
	r.HandleFunc(ApiPrefix+"/series/{id}", controllers.GetSeries).Methods(http.MethodGet)

	// Titles routes
	r.HandleFunc(ApiPrefix+"/titles", controllers.CreateTitle).Methods(http.MethodPost)
	r.HandleFunc(ApiPrefix+"/titles/{id}", controllers.GetTitle).Methods(http.MethodGet)

	// Users routes
	r.HandleFunc(ApiPrefix+"/users/{id}", controllers.GetCurrentUser).Methods(http.MethodGet)
	//r.HandleFunc(ApiPrefix+"/users/{id}", controllers.GetUser).Methods(http.MethodGet)

	// Service routes
	//r.HandleFunc(ApiPrefix + "/notify/{id}", controllers.Notify).Methods(http.MethodGet)

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
