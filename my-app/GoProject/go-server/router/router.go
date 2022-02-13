package router

import (
	"go-server/middleware"

	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	// Init the mux router
	router := mux.NewRouter()

	// Get all actors
	router.HandleFunc("/actors/", middleware.GetActors).Methods("GET")
	// Get all artists
	router.HandleFunc("/artists/", middleware.GetArtists).Methods("GET")
	// Get all customers
	router.HandleFunc("/customers/", middleware.GetCustomers).Methods("GET")
	// Get all employees
	router.HandleFunc("/employees/", middleware.GetEmployees).Methods("GET")
	return router
}
