package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/RomanoLu/events-to-go/src/goevent/model"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type result struct {
	Success string `json:"success"`
}

func sendJson(w http.ResponseWriter, value interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(value); err != nil {
		log.Errorf("Failure encoding value to JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getId(r *http.Request) (uint, error) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 0)
	if err != nil {
		log.Errorf("Can't get ID from request: %v", err)
		return 0, err
	}
	fmt.Println(id)
	return uint(id), nil
}

func getEventIdAndUserId(r *http.Request) (uint,uint, error) {
	vars := mux.Vars(r)
	eventid, err := strconv.ParseUint(vars["id"], 10, 0)
	if err != nil {
		log.Errorf("Can't get ID from request: %v", err)
		return 0, 0, err
	}
	fmt.Println(eventid)
	userid, err := strconv.ParseUint(vars["userid"], 10, 0)
	if err != nil {
		log.Errorf("Can't get ID from request: %v", err)
		return 0, 0, err
	}
	fmt.Println(userid)
	return uint(eventid),uint(userid), nil
}

func getGeodata(r *http.Request) (float64,float64, error) {
	vars := mux.Vars(r)
	long, err := strconv.ParseFloat(vars["longitude"], 64)
	if err != nil {
		log.Errorf("Can't get Longitude from request: %v", err)
		return 0, 0, err
	}
	lat, err := strconv.ParseFloat(vars["latitude"], 64)
	if err != nil {
		log.Errorf("Can't get Latitude from request: %v", err)
		return 0,0, err
	}
	log.Infoln("Recived %v as Longitude and %v as Latitude",long, lat)
	return long, lat, nil
}

func getIds(r *http.Request) (uint,uint, model.Invetation, error) {
	var invetation model.Invetation
	err := json.NewDecoder(r.Body).Decode(&invetation)
	if err != nil {
		log.Errorf("Can't serialize request body to event struct: %v", err)
		return 0,0,invetation ,err
	}
	log.Info("Invetation message: %v", invetation.Message)
	vars := mux.Vars(r)
	eventid, err := strconv.ParseUint(vars["eventid"], 10, 0)
	if err != nil {
		log.Errorf("Can't get ID from request: %v", err)
		return 0,0,invetation,err
	}
	
	userid, err := strconv.ParseUint(vars["userid"], 10, 0)
	if err != nil {
		log.Errorf("Can't get ID from request: %v", err)
		return 0,0,invetation,err
	}
	log.Infoln("Recived %v as eventid and %v as userid",eventid, userid)
	return uint(eventid), uint(userid),invetation, nil
}
