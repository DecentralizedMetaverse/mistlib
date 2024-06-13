package world_creation

import (
	"encoding/json"
	"net/http"
)

type GetComponentRequest struct {
	ObjectID    string `json:"object_id"`
	ComponentID string `json:"component_id"`
}

type GetComponentResponse struct {
	Data string `json:"data"`
}

func GetComponent(w http.ResponseWriter, r *http.Request) {
	var req GetComponentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := GetComponentResponse{
		Data: "Example data",
	}
	json.NewEncoder(w).Encode(res)
}
