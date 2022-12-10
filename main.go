package main

import (
	"fmt"
	"log"
	"net/http"

	"api-server/controllers"
	"api-server/database"

	"github.com/gorilla/mux"
)

func main() {
	const port = 5000

	// Load Database - MongoDB
	database.Init()

	// Initialize the router
	router := mux.NewRouter().StrictSlash(true)

	// Register Routes
	RegisterProductRoutes(router)

	// Start the server
	log.Println(fmt.Sprintln("Starting Server on port", port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), router))
}

func RegisterProductRoutes(router *mux.Router) {
	router.HandleFunc("/api/get", controllers.GetAll).Methods("POST")
	router.HandleFunc("/api/create/{id}", controllers.Create).Methods("POST")
	router.HandleFunc("/api/update/{id}", controllers.Update).Methods("POST")
	router.HandleFunc("/api/delete/{id}", controllers.Delete).Methods("POST")
}
