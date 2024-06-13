package basic

import (
	"encoding/json"
	"net/http"
)

type DisconnectRequest struct {
	UserID string `json:"user_id"`
}

func Disconnect(w http.ResponseWriter, r *http.Request) {
	var req DisconnectRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
