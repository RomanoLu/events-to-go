package handler

import (
	"encoding/json"
	"net/http"

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

func GetEventByDate(w http.ResponseWriter, _ *http.Request) {

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
		log.Errorf("Failure updating event with ID %v: %v", id, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if event == nil {
		http.Error(w, "404 event not found", http.StatusNotFound)
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

func InvideUser(w http.ResponseWriter, r *http.Request) {

	/* Hier muss jetzt noch die Invetation dem User zugeordnet werden.
	Dafür kann die unten erstellte invetation einfach dem user mit der oben stehenden UserID hinzhugefügt wederden
	Dafür muss dann eine service methode geschrieben werden die das macht
	+ noch irgendwas überlegen wie man ein event besser einer invetation zuordnen kann
	*/
	eventID, _, message, _ := getIds(r)

	invetation := model.Invetation{
		Message:  message,
		Accepted: false,
		Eventid:  eventID,
	}

	sendJson(w, invetation)
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
