package service

import (
	"time"

	"github.com/RomanoLu/events-to-go/src/goevent/db"
	"github.com/RomanoLu/events-to-go/src/goevent/model"
	log "github.com/sirupsen/logrus"
)

func CreateEvent(event *model.Event) error {
	result := db.DB.Create(event)
	if result.Error != nil {
		return result.Error
	}
	log.Infof("Successfully stored new event with ID %v in database.", event.ID)
	log.Tracef("Stored: %v", event)
	return nil
}

func GetAllEvents() ([]model.Event, error){
	var events []model.Event
	result := db.DB.Find(&events)
	if result.Error != nil {
		return nil, result.Error
	}
	log.Tracef("Retrieved: %v", events)
	return events, nil
}

func GetEventById(id uint) (*model.Event, error) {
	var event *model.Event
	result := db.DB.Find(&event, id)
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
	db.DB.Model(&existingevent).Updates(event)
	entry := log.WithField("ID", id)
	entry.Info("Successfully updated event.")
	entry.Tracef("Updated: %v", event)
	return existingevent, nil
}

func DeleteEvent(id uint) (*model.Event, error) {
	event, err := GetEventById(id)
	if err != nil {
		log.Error("Kein Event gefunden")
		return event, err
	}
	db.DB.Delete(event)
	entry := log.WithField("ID", id)
	entry.Info("Successfully deleted event.")
	entry.Tracef("Deleted: %v", event)
	return event, nil
}
