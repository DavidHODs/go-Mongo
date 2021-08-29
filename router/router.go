package router

import (
	"net/http"

	"github.com/DavidHODs/go-Mongo/handlers"
	"github.com/gorilla/mux"
)


func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/person", handlers.CreatePersonEndPoint).Methods("POST")
	router.HandleFunc("/people", handlers.GetPeopleEndPoint).Methods("GET")
	router.HandleFunc("/person/{id}", handlers.GetPersonEndPoint).Methods("GET")

	http.Handle("/", router)

	return router
}
