package communication

import (
	"encoding/json"
	"net/http"
)

type SendGroupMessageRequest struct {
	GroupID string `json:"group_id"`
	UserID  string `json:"user_id"`
	Message string `json:"message"`
}

func SendGroupMessage(w http.ResponseWriter, r *http.Request) {
	var req SendGroupMessageRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
