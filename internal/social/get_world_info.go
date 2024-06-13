package social

import (
	"encoding/json"
	"net/http"
)

type GetWorldInfoRequest struct {
	WorldID string `json:"world_id"`
}

type GetWorldInfoResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func GetWorldInfo(w http.ResponseWriter, r *http.Request) {
	var req GetWorldInfoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := GetWorldInfoResponse{
		Name:        "Example World",
		Description: "This is an example world.",
	}
	json.NewEncoder(w).Encode(res)
}
