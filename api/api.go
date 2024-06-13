package api

import (
	"mistlib/internal/basic"
	"mistlib/internal/communication"
	"mistlib/internal/programming"
	"mistlib/internal/social"
	"mistlib/internal/world_creation"
	"net/http"

	"github.com/gorilla/mux"
)

func StartServer() {
	r := mux.NewRouter()

	// Basic endpoints
	r.HandleFunc("/connect", basic.Connect).Methods("POST")
	r.HandleFunc("/disconnect", basic.Disconnect).Methods("POST")
	r.HandleFunc("/joinworld", basic.JoinWorld).Methods("POST")
	r.HandleFunc("/leaveworld", basic.LeaveWorld).Methods("POST")
	r.HandleFunc("/sendlocation", basic.SendLocation).Methods("POST")
	r.HandleFunc("/sendanimation", basic.SendAnimation).Methods("POST")

	// Social endpoints
	r.HandleFunc("/addfriend", social.AddFriend).Methods("POST")
	r.HandleFunc("/removefriend", social.RemoveFriend).Methods("POST")
	r.HandleFunc("/getfriendlist", social.GetFriendList).Methods("POST")
	r.HandleFunc("/getuserlist", social.GetUserList).Methods("POST")
	r.HandleFunc("/getworldlist", social.GetWorldList).Methods("POST")
	r.HandleFunc("/getworldinfo", social.GetWorldInfo).Methods("POST")

	// Communication endpoints
	r.HandleFunc("/sendmessage", communication.SendMessage).Methods("POST")
	r.HandleFunc("/sendwhisper", communication.SendWhisper).Methods("POST")
	r.HandleFunc("/creategroup", communication.CreateGroup).Methods("POST")
	r.HandleFunc("/joingroup", communication.JoinGroup).Methods("POST")
	r.HandleFunc("/leavegroup", communication.LeaveGroup).Methods("POST")
	r.HandleFunc("/sendgroupmessage", communication.SendGroupMessage).Methods("POST")

	// World Creation endpoints
	r.HandleFunc("/putobject", world_creation.PutObject).Methods("POST")
	r.HandleFunc("/addcomponent", world_creation.AddComponent).Methods("POST")
	r.HandleFunc("/removecomponent", world_creation.RemoveComponent).Methods("POST")
	r.HandleFunc("/getcomponent", world_creation.GetComponent).Methods("POST")
	r.HandleFunc("/getcomponents", world_creation.GetComponents).Methods("POST")

	// Programming endpoints
	r.HandleFunc("/rpc", programming.RPCHandler).Methods("POST")

	http.ListenAndServe(":8080", r)
}
