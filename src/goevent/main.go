package main

import (
	"net/http"

	"github.com/RomanoLu/events-to-go/src/goevent/db"
	"github.com/RomanoLu/events-to-go/src/goevent/handler"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Println("Starting My-Aktion API server")
	router := mux.NewRouter()

	//Event Requests
	router.HandleFunc("/event", handler.CreateEvent).Methods("POST")
	router.HandleFunc("/event", handler.GetAllEvents).Methods("GET")
	router.HandleFunc("/event/{id}", handler.GetEventById).Methods("GET")
	router.HandleFunc("/event/{longitude}/{latitude}", handler.GetEventsNearby).Methods("GET")
	router.HandleFunc("/event/{id}", handler.UpdateEvent).Methods("UPDATE")
	router.HandleFunc("/event/{id}", handler.DeleteEvent).Methods("DELETE")	
	router.HandleFunc("/event/invide/{eventid}/{userid}", handler.InvideUser).Methods("POST")
	router.HandleFunc("/location", handler.GetLocations).Methods("GET")

	//User requests	
	router.HandleFunc("/user",handler.GetAllUsers).Methods("GET")
	router.HandleFunc("/user",handler.CreateUser).Methods("POST")
	router.HandleFunc("/user/{id}",handler.GetUserById).Methods("GET")
	router.HandleFunc("/user/invetation/{id}",handler.GetInvetations).Methods("GET")




	
	//router.HandleFunc("/location",handler.CreateLocation).Methods("POST")
	//router.HandleFunc("/event/{userid}",handler.AddUserToEvent).Methods("POST")
	//router.HandleFunc("/event/{location}",handler.DeleteEvent).Methods("DELETE")
	//router.HandleFunc("/event/{Date}",handler.DeleteEvent).Methods("DELETE")
	//router.HandleFunc("/acceptEvent/{userid}",handler.AcceptEvent).Methods("POST")
	//router.HandleFunc("/invetation/",handler.CreateInvetation).Methods("POST")
	//router.HandleFunc("/invetations/{userid}",handler.GetInvetation).Methods("GET")

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
