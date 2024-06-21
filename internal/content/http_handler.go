package content

import (
	"fmt"
	"net/http"
	"strings"
)

func HandleInitAPI(w http.ResponseWriter, r *http.Request) {
	handleInit(nil)
	w.Write([]byte("Repository initialized."))
}

func HandleSwitchAPI(w http.ResponseWriter, r *http.Request) {
	worldName := r.URL.Query().Get("worldName")
	handleSwitch([]string{worldName})
	w.Write([]byte(fmt.Sprintf("Switched to world: %s", worldName)))
}

func HandleGetAPI(w http.ResponseWriter, r *http.Request) {
	args := r.URL.Query().Get("args")
	handleGet(strings.Split(args, " "))
	w.Write([]byte("Get operation completed."))
}

func HandleAddAPI(w http.ResponseWriter, r *http.Request) {
	args := r.URL.Query().Get("args")
	handleAdd(strings.Split(args, " "))
	w.Write([]byte("Put operation completed."))
}

func HandleSetPasswordAPI(w http.ResponseWriter, r *http.Request) {
	password := r.URL.Query().Get("password")
	handleSetPassword([]string{password})
	w.Write([]byte("Password set successfully."))
}

func HandleCatAPI(w http.ResponseWriter, r *http.Request) {
	fileHash := r.URL.Query().Get("fileHash")
	handleCat([]string{fileHash})
	w.Write([]byte("Cat operation completed."))
}

func HandleGetWorldCIDAPI(w http.ResponseWriter, r *http.Request) {
	handleGetWorldCID(nil)
	w.Write([]byte("Get World CID operation completed."))
}

func HandleDownloadWorldAPI(w http.ResponseWriter, r *http.Request) {
	cid := r.URL.Query().Get("cid")
	handleDownloadWorld([]string{cid})
	w.Write([]byte(fmt.Sprintf("Download world operation completed for CID: %s", cid)))
}

func HandleGetWorldInfoAPI(w http.ResponseWriter, r *http.Request) {
	args := r.URL.Query().Get("args")
	handleGetWorldInfo(strings.Split(args, " "))
	w.Write([]byte("Get world data operation completed."))
}

func HandleSetCustomDataAPI(w http.ResponseWriter, r *http.Request) {
	args := r.URL.Query().Get("args")
	handleSetCustomData(strings.Split(args, " "))
	w.Write([]byte("Set custom data operation completed."))
}

func HandleSetParentAPI(w http.ResponseWriter, r *http.Request) {
	args := r.URL.Query().Get("args")
	handleSetParent(strings.Split(args, " "))
	w.Write([]byte("Set parent operation completed."))
}

func HandleUpdateAPI(w http.ResponseWriter, r *http.Request) {
	args := r.URL.Query().Get("args")
	handleUpdate(strings.Split(args, " "))
	w.Write([]byte("Update operation completed."))
}
