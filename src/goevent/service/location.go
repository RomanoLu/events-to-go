package service

import (
	"github.com/RomanoLu/events-to-go/src/goevent/db"
	"github.com/RomanoLu/events-to-go/src/goevent/model"
	log "github.com/sirupsen/logrus"
)

func GetLocations() ([]model.Location, error){
	var locations []model.Location
	result := db.DB.Find(&locations)
	if result.Error != nil {
		return nil, result.Error
	}
	log.Tracef("Retrieved: %v", locations)
	return locations, nil
}

