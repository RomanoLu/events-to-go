package handler

import (
	"encoding/json"
	"net/http"

	"github.com/RomanoLu/events-to-go/src/goevent/model"
	"github.com/RomanoLu/events-to-go/src/goevent/service"
	log "github.com/sirupsen/logrus"
)


func CreateUser(w http.ResponseWriter, r *http.Request) {
	user, err := getUser(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := service.CreateUser(user); err != nil {
		log.Errorf("Error calling service CreateCampaign: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendJson(w, user)
}

func GetUserById(w http.ResponseWriter, r *http.Request){
	id, err := getId(r)
	log.Trace("Die id ist: %v", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := service.GetUserById(id)
	if err != nil {
		log.Errorf("Failure retrieving user with ID %v: %v", id, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if user == nil {
		http.Error(w, "404 user not found", http.StatusNotFound)
		return
	}
	sendJson(w, user)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request){
	user, err := service.GetAllUsers()
	if err != nil {
		log.Errorf("Error calling service GetAllUsers: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sendJson(w, user)
}

func getUser(r *http.Request) (*model.User, error) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Errorf("Can't serialize request body to event struct: %v", err)
		return nil, err
	}
	return &user, nil
}