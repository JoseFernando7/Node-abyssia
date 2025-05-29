package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/josefernando7/node-abyssia/internal/containers"
	"github.com/josefernando7/node-abyssia/internal/handlers"
)

func main() {
	manager := containers.NewDockerManager()

	r := mux.NewRouter()

	r.HandleFunc("/containers", func(w http.ResponseWriter, r *http.Request) {
		list, err := manager.ListContainers()
		if err != nil {
			errorMessage := fmt.Sprintf("Error listing containers: %v", err)
			http.Error(w, errorMessage, http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(list)
	}).Methods("GET")

	r.HandleFunc("/containers/create", handlers.CreateContainerHandler).Methods("POST")

	r.HandleFunc("/containers/delete/{id}", handlers.DeleteContainerHandler).Methods("DELETE")

	log.Println("✨ Node†abyssia corriendo en http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
