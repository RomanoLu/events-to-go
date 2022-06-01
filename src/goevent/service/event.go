package service

import (
	"github.com/RomanoLu/events-to-go/src/goevent/db"
	"github.com/RomanoLu/events-to-go/src/goevent/model"
	log "github.com/sirupsen/logrus"
)

func AddEvent(event *model.Event) error {
	result := db.DB.Create(event)
	if result.Error != nil {
		return result.Error
	}
	log.Infof("Successfully stored new event with ID %v in database.", event.ID)
	log.Tracef("Stored: %v", event)
	return nil
}