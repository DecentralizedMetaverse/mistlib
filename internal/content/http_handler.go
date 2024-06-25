package content

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type APIResponse struct {
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

func sendJSONResponse(w http.ResponseWriter, status int, message string, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	response := APIResponse{Message: message}
	if err != nil {
		response.Error = err.Error()
	}
	json.NewEncoder(w).Encode(response)
}

func HandleInitAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		sendJSONResponse(w, http.StatusMethodNotAllowed, "", fmt.Errorf("method not allowed"))
		return
	}
	handleInit(nil)
	sendJSONResponse(w, http.StatusOK, "Repository initialized.", nil)
}

func HandleSwitchAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		sendJSONResponse(w, http.StatusMethodNotAllowed, "", fmt.Errorf("method not allowed"))
		return
	}
	args := r.URL.Query().Get("args")
	handleSwitch(strings.Split(args, " "))
	sendJSONResponse(w, http.StatusOK, "Switch operation completed.", nil)
}

func HandleGetAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		sendJSONResponse(w, http.StatusMethodNotAllowed, "", fmt.Errorf("method not allowed"))
		return
	}
	args := r.URL.Query().Get("args")
	handleGet(strings.Split(args, " "))
	sendJSONResponse(w, http.StatusOK, "Get operation completed.", nil)
}

func HandleAddAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		sendJSONResponse(w, http.StatusMethodNotAllowed, "", fmt.Errorf("method not allowed"))
		return
	}
	args := r.URL.Query().Get("args")
	handleAdd(strings.Split(args, " "))
	sendJSONResponse(w, http.StatusOK, "Add operation completed.", nil)
}

func HandleRemoveAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		sendJSONResponse(w, http.StatusMethodNotAllowed, "", fmt.Errorf("method not allowed"))
		return
	}
	args := r.URL.Query().Get("args")
	handleRemove(strings.Split(args, " "))
	sendJSONResponse(w, http.StatusOK, "Remove operation completed.", nil)
}

func HandleSetPasswordAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		sendJSONResponse(w, http.StatusMethodNotAllowed, "", fmt.Errorf("method not allowed"))
		return
	}
	password := r.URL.Query().Get("password")
	handleSetPassword([]string{password})
	sendJSONResponse(w, http.StatusOK, "Password set successfully.", nil)
}

func HandleCatAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		sendJSONResponse(w, http.StatusMethodNotAllowed, "", fmt.Errorf("method not allowed"))
		return
	}
	fileHash := r.URL.Query().Get("fileHash")
	handleCat([]string{fileHash})
	sendJSONResponse(w, http.StatusOK, "Cat operation completed.", nil)
}

func HandleGetWorldCIDAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		sendJSONResponse(w, http.StatusMethodNotAllowed, "", fmt.Errorf("method not allowed"))
		return
	}
	handleGetWorldCID(nil)
	sendJSONResponse(w, http.StatusOK, "Get World CID operation completed.", nil)
}

func HandleDownloadWorldAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		sendJSONResponse(w, http.StatusMethodNotAllowed, "", fmt.Errorf("method not allowed"))
		return
	}
	cid := r.URL.Query().Get("cid")
	handleDownloadWorld([]string{cid})
	sendJSONResponse(w, http.StatusOK, fmt.Sprintf("Download world operation completed for CID: %s", cid), nil)
}

func HandleGetWorldInfoAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		sendJSONResponse(w, http.StatusMethodNotAllowed, "", fmt.Errorf("method not allowed"))
		return
	}
	args := r.URL.Query().Get("args")
	handleGetWorldInfo(strings.Split(args, " "))
	sendJSONResponse(w, http.StatusOK, "Get world data operation completed.", nil)
}

func HandleUnpackAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		sendJSONResponse(w, http.StatusMethodNotAllowed, "", fmt.Errorf("method not allowed"))
		return
	}
	args := r.URL.Query().Get("args")
	handleUnpack(strings.Split(args, " "))
	sendJSONResponse(w, http.StatusOK, "Unpack operation completed.", nil)
}

func HandleSetCustomDataAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		sendJSONResponse(w, http.StatusMethodNotAllowed, "", fmt.Errorf("method not allowed"))
		return
	}
	args := r.URL.Query().Get("args")
	handleSetCustomData(strings.Split(args, " "))
	sendJSONResponse(w, http.StatusOK, "Set custom data operation completed.", nil)
}

func HandleSetParentAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		sendJSONResponse(w, http.StatusMethodNotAllowed, "", fmt.Errorf("method not allowed"))
		return
	}
	args := r.URL.Query().Get("args")
	handleSetParent(strings.Split(args, " "))
	sendJSONResponse(w, http.StatusOK, "Set parent operation completed.", nil)
}

func HandleUpdateAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		sendJSONResponse(w, http.StatusMethodNotAllowed, "", fmt.Errorf("method not allowed"))
		return
	}
	args := r.URL.Query().Get("args")
	handleUpdate(strings.Split(args, " "))
	sendJSONResponse(w, http.StatusOK, "Update operation completed.", nil)
}
