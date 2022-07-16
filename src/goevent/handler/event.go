package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/RomanoLu/events-to-go/src/goevent/model"
	"github.com/RomanoLu/events-to-go/src/goevent/service"
	log "github.com/sirupsen/logrus"
)

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	event, err := getEvent(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := service.CreateEvent(event); err != nil {
		log.Errorf("Error calling service CreateCampaign: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendJson(w, event)
}

func GetAllEvents(w http.ResponseWriter, _ *http.Request) {
	event, err := service.GetAllEvents()
	if err != nil {
		log.Errorf("Error calling service GetAllEvents: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sendJson(w, event)
}

func GetEventById(w http.ResponseWriter, r *http.Request) {
	id, err := getId(r)
	log.Trace("Die id ist: %v", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	event, err := service.GetEventById(id)
	if err != nil {
		log.Errorf("Failure retrieving event with ID %v: %v", id, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if event == nil {
		http.Error(w, "404 event not found", http.StatusNotFound)
		return
	}
	sendJson(w, event)
}

func GetParticipants(w http.ResponseWriter, r *http.Request){
	id, err := getId(r)
	log.Trace("Die id ist: %v", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	participants, err := service.GetParticipants(id)
	if err != nil {
		log.Errorf("Failure retrieving Participants from event with ID %v: %v", id, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var  user []model.User= participants.Participants 
	log.Info("Recieved Participants: %v", user)
	sendJson(w, user)
}
func AddParticipants(w http.ResponseWriter, r *http.Request){
	eventid, userid, err := getEventIdAndUserId(r)
	log.Trace("Die id ist: %v & %v", eventid, userid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	user, err := service.GetUserById(userid)
	log.Info("Dieser User sollte hinzugefügt werden: %v", user)
	event, err := service.AddParticipants(eventid, user)
	if err != nil {
		log.Error("Error trying to add Participant")
		return
	}
	sendJson(w, event)
}

func GetEventsNearby(w http.ResponseWriter, r *http.Request) {
	long, lat, err := getGeodata(r)
	log.Trace("Die Location id ist: %v ; %v", long, lat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	event, err := service.GetEventByLocation(long, lat)
	if err != nil {
		log.Errorf("Failure retrieving Longitude %v and latitude %v", long, lat, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if event == nil {
		http.Error(w, "Leider gibt es keine Events in deiner Nähe", http.StatusNotFound)
		return
	}
	sendJson(w, event)
}

func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	id, err := getId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	event, err := getEvent(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	event, err = service.UpdateEvent(id, event)
	if err != nil {
		log.Errorf("Failure updating campaign with ID %v: %v", id, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if event == nil {
		http.Error(w, "404 campaign not found", http.StatusNotFound)
		return
	}
	sendJson(w, event)


}


func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	id, err := getId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	event, err := service.DeleteEvent(id)
	if err != nil {
		log.Errorf("Failure deleting event with ID %v: %v", id, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if event == nil {
		http.Error(w, "404 event not found", http.StatusNotFound)
		return
	}
	sendJson(w, result{Success: "OK"})
}
/*
func InvideUser(w http.ResponseWriter, r *http.Request) {
	_, userid, invetation, _ := getIds(r)
	log.Info("hier die invetation:  %v", invetation)
	user, _ :=service.CreateInvetation(userid, invetation)

	sendJson(w, user)
}
*/
func getEvent(r *http.Request) (*model.Event, error) {
	var event model.Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		log.Errorf("Can't serialize request body to event struct: %v", err)
		return nil, err
	}
	return &event, nil
}


func GetEventByDate(w http.ResponseWriter, _ *http.Request) {
	currentTime := time.Now()
	events, err := service.GetEventByDate(currentTime)
	if err != nil {
		log.Errorf("Failure deleting event with ID %v: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendJson(w, events)
}


func GetNextWeekendEvents(w http.ResponseWriter, _ *http.Request){
	currentTime := time.Now()
	events, err := service.GetNextWeekendEvents(currentTime)
	if err != nil {
		log.Errorf("Failure deleting event with ID %v: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendJson(w, events)
}

func SaveEventInCalendar(w http.ResponseWriter, r *http.Request){
	id, err := getId(r)
	if err!= nil {
		return
	}
	_, s, err := service.SaveEventInCalendar(id)
	if err != nil {
		log.Errorf("Failure creating calendar event with ID %v: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendJson(w, s)
}