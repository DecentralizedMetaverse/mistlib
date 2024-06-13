package basic

import (
	"encoding/json"
	"net/http"
)

type Location struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

type SendLocationRequest struct {
	UserID   string   `json:"user_id"`
	Location Location `json:"location"`
}

func SendLocation(w http.ResponseWriter, r *http.Request) {
	var req SendLocationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
