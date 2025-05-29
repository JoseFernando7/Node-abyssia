package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/josefernando7/node-abyssia/internal/containers"
)

func DeleteContainerHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	containerID := vars["id"]

	fmt.Println("Attempting to delete container with ID:", containerID)

	if err := containers.DeleteContainer(containerID); err != nil {
		http.Error(w, "Failed to delete container: " + err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"message": "Container deleted successfully",
		"container_id": containerID,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
