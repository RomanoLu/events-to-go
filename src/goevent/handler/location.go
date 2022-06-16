package handler

import (
	"net/http"

	"github.com/RomanoLu/events-to-go/src/goevent/service"
	log "github.com/sirupsen/logrus"
)


func GetLocations(w http.ResponseWriter, _ *http.Request)  {
	log.Info("Methode wird aufgerufen")
	locations, err := service.GetLocations()
	if err != nil {
		log.Errorf("Error calling service GetAllEvents: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sendJson(w, locations)
}