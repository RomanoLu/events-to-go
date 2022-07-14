package handler

import (
	"encoding/json"
	"net/http"

	"github.com/RomanoLu/events-to-go/src/goevent/model"
	"github.com/RomanoLu/events-to-go/src/goevent/service"
	log "github.com/sirupsen/logrus"
)

func CreateInventation(w http.ResponseWriter, r *http.Request) {
	invetation, err := getInvetation(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := service.CreateInventation(invetation); err != nil {
		log.Errorf("Error calling service CreateInventation: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendJson(w, invetation)
}
func GetInvetations(w http.ResponseWriter, r *http.Request) {
	id, err := getId(r)
	log.Trace("Die id ist: %v", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	invetation, err := service.GetInvetationByUserID(id)
	if err != nil {
		log.Errorf("Error calling service GetInvetations: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sendJson(w, invetation)
}

func AcceptInvetation(w http.ResponseWriter, r *http.Request) {
	id, err := getId(r)
	log.Trace("Die id ist: %v", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	invetation, event, err := service.AcceptInvetation(id)
	if err != nil {
		log.Errorf("Error calling service GetInvetations: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sendJson(w, invetation)
	sendJson(w, event)
}

func getInvetation(r *http.Request) (*model.Invetation, error) {
	var invetation model.Invetation
	err := json.NewDecoder(r.Body).Decode(&invetation)
	if err != nil {
		log.Errorf("Can't serialize request body to event struct: %v", err)
		return nil, err
	}
	return &invetation, nil
}
