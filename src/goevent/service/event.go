package service

import (
	"math"
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

func GetAllEvents() ([]model.Event, error) {
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

func GetEventByLocation(longitude float64, latitude float64) ([]model.Event, error) {
	events, _ := GetAllEvents()
	var nearbyEvents []model.Event

	for _, e2 := range events {
		loc, err := getLocationByEvent(e2)
		if err != nil {
			log.Errorf("Failure retrieving Location")
			return nil, err
		}
		if distance(loc.Latitude, loc.Longitude, latitude, longitude) < 20 {
			nearbyEvents = append(nearbyEvents, e2)
		}
	}

	return nearbyEvents, nil
}

func getLocationByEvent(event model.Event) (*model.Location, error) {
	var location *model.Location
	result := db.DB.Find(&location, event.LocationID)
	if result.Error != nil {
		return nil, result.Error
	}
	log.Tracef("Retrieved: %v", event)
	return location, nil
}

func GetEventByDate(date time.Time) (*model.Event, error) {
	return nil, nil
}

func UpdateEvent(id uint, event *model.Event) (*model.Event, error) {
	existingEvent, err := GetEventById(id)
	if existingEvent == nil || err != nil {
		return existingEvent, err
	}
	existingEvent.Begin = event.Begin
	existingEvent.Description = event.Description
	existingEvent.End = event.End
	existingEvent.Titel = event.Titel
	existingEvent.MaxNumberOfParticipants = event.MaxNumberOfParticipants
	existingEvent.Participants = event.Participants
	existingEvent.LocationID = event.LocationID
	existingEvent.Type = event.Type
	existingEvent.Host = event.Host

	result := db.DB.Save(existingEvent)
	if result.Error != nil {
		return nil, result.Error
	}
	entry := log.WithField("ID", id)
	entry.Info("Successfully updated event.")
	entry.Tracef("Updated: %v", existingEvent)
	return existingEvent, nil

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

func distance(lat1 float64, lng1 float64, lat2 float64, lng2 float64, unit ...string) float64 {
	const PI float64 = 3.141592653589793

	radlat1 := float64(PI * lat1 / 180)
	radlat2 := float64(PI * lat2 / 180)

	theta := float64(lng1 - lng2)
	radtheta := float64(PI * theta / 180)

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)

	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / PI
	dist = dist * 60 * 1.1515

	if len(unit) > 0 {
		if unit[0] == "K" {
			dist = dist * 1.609344
		} else if unit[0] == "N" {
			dist = dist * 0.8684
		}
	}

	return dist
}
