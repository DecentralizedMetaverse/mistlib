package basic

import (
	"encoding/json"
	"net/http"
)

type ConnectRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Connect(w http.ResponseWriter, r *http.Request) {
	var req ConnectRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
