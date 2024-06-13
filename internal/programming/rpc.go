package programming

import (
	"encoding/json"
	"net/http"
)

type RPCRequest struct {
	Method string `json:"method"`
	Params string `json:"params"`
}

type RPCResponse struct {
	Result string `json:"result"`
}

func RPCHandler(w http.ResponseWriter, r *http.Request) {
	var req RPCRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := RPCResponse{
		Result: "Example result",
	}
	json.NewEncoder(w).Encode(res)
}
