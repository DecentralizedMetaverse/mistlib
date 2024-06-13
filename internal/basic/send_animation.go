package basic

import (
	"encoding/json"
	"net/http"
)

type SendAnimationRequest struct {
	UserID    string `json:"user_id"`
	Animation string `json:"animation"`
}

func SendAnimation(w http.ResponseWriter, r *http.Request) {
	var req SendAnimationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
