package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/josefernando7/node-abyssia/internal/containers"
	"github.com/josefernando7/node-abyssia/pkg/models"
)

func CreateContainerHandler(w http.ResponseWriter, r *http.Request) {
	var request models.ContainerCreateRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	id, err := containers.CreateContainer(request)
	if err != nil {
		http.Error(w, "Failed to create container: " + err.Error(), http.StatusInternalServerError)
	}

	response := map[string]string {
		"message": "Container created successfully",
		"container_id": id,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
