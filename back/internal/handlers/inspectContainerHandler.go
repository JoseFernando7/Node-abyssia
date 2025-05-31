package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/josefernando7/node-abyssia/internal/containers"
)

func InspectContainerHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	containerID := vars["id"]

	inspectResponse, err := containers.InspectContainer(containerID)
	if err != nil {
		http.Error(w, "Failed to inspect container: " + err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(inspectResponse)
}
