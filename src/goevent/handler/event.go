package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/RomanoLu/events-to-go/src/goevent/db"
	"github.com/RomanoLu/events-to-go/src/goevent/model"
	"github.com/RomanoLu/events-to-go/src/goevent/service"
	log "github.com/sirupsen/logrus"
)

func AddEvent(w http.ResponseWriter, r *http.Request) {
	event, err := getEvent(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := service.AddEvent(event); err != nil {
		log.Errorf("Error calling service CreateCampaign: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendJson(w, event)
}

func GetAllEvents(w http.ResponseWriter, _ *http.Request)  {
	var events []model.Event
	result := db.DB.Preload("User").Find(&events)
	if result.Error != nil {
		return
	}
	log.Tracef("Retrieved: %v", events)
	sendJson(w, events)
}

func GetEventById(id uint) (*model.Event, error) {
	var event *model.Event
	result := db.DB.Preload("User").Find(&event, id)
	if result.Error != nil {
		return nil, result.Error
	}
	log.Tracef("Retrieved: %v", event)
	return event, nil
}

func GetEventByLocation(location model.Location) (*model.Event, error) {
	return nil, nil
}

func GetEventByDate(date time.Time) (*model.Event, error) {
	return nil, nil
}

func UpdateEvent(id uint, event *model.Event) (*model.Event, error) {
	existingevent, err := GetEventById(id)
	if err != nil {
		return existingevent, err
	}
	db.DB.Preload("User").Model(&existingevent).Updates(event)
	entry := log.WithField("ID", id)
	entry.Info("Successfully updated event.")
	entry.Tracef("Updated: %v", event)
	return existingevent, nil
}
func DeleteEvent(id uint) (*model.Event, error) {
	event, err := GetEventById(id)
	if err != nil {
		return event, err
	}
	db.DB.Delete(event)
	entry := log.WithField("ID", id)
	entry.Info("Successfully deleted event.")
	entry.Tracef("Deleted: %v", event)
	return event, nil
}

func getEvent(r *http.Request) (*model.Event, error) {
	var event model.Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		log.Errorf("Can't serialize request body to event struct: %v", err)
		return nil, err
	}
	return &event, nil
}