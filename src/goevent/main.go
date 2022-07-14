package main

import (
	"net/http"

	"github.com/RomanoLu/events-to-go/src/goevent/db"
	"github.com/RomanoLu/events-to-go/src/goevent/handler"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Println("Starting Events-to-go API server")
	router := mux.NewRouter()

	//Event Requests
	router.HandleFunc("/event", handler.CreateEvent).Methods("POST")
	router.HandleFunc("/event", handler.GetAllEvents).Methods("GET")
	router.HandleFunc("/event/{id}", handler.GetEventById).Methods("GET")
	router.HandleFunc("/event/{longitude}/{latitude}", handler.GetEventsNearby).Methods("GET")
	router.HandleFunc("/eventsoftoday", handler.GetEventByDate).Methods("GET")
	router.HandleFunc("/eventsofweekend", handler.GetNextWeekendEvents).Methods("GET")
	router.HandleFunc("/event/{id}", handler.UpdateEvent).Methods("UPDATE")
	router.HandleFunc("/event/{id}", handler.DeleteEvent).Methods("DELETE")	
	router.HandleFunc("/event/invide", handler.CreateInventation).Methods("POST")
	router.HandleFunc("/location", handler.GetLocations).Methods("GET")
	router.HandleFunc("/adduser/{eventid}/{userid}",handler.AddParticipants).Methods("POST")

	//User requests	
	router.HandleFunc("/user",handler.GetAllUsers).Methods("GET")
	router.HandleFunc("/user",handler.CreateUser).Methods("POST")
	router.HandleFunc("/user/{id}",handler.GetUserById).Methods("GET")
	router.HandleFunc("/user/invetation/{id}",handler.GetInvetations).Methods("GET")
	router.HandleFunc("/acceptinvetation/{id}",handler.AcceptInvetation).Methods("POST")
	router.HandleFunc("/participants/{id}",handler.GetParticipants).Methods("POST")
	//Save Event in Calender
	
	//router.HandleFunc("/event/{location}",handler.DeleteEvent).Methods("DELETE")
	//router.HandleFunc("/event/{Date}",handler.DeleteEvent).Methods("DELETE")


	if err := http.ListenAndServe(":8000", router); err != nil {
		log.Fatal(err)
	}
}

func init() {
	db.Init()
	// init logger
	log.SetLevel(log.TraceLevel)
	log.SetFormatter(&log.TextFormatter{})
	log.SetReportCaller(true)
	/*level, err := log.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		log.Info("Log level not specified, set default to: INFO")
		log.SetLevel(log.InfoLevel)
		return
	}
	log.SetLevel(level)*/
}
