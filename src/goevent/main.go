package main

import (
	"net/http"

	"github.com/RomanoLu/events-to-go/src/goevent/db"
	"github.com/RomanoLu/events-to-go/src/goevent/handler"
	"github.com/RomanoLu/events-to-go/src/goevent/middleware"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Println("Starting Events-to-go API server")
	router := mux.NewRouter()

	router.HandleFunc("/jwt", middleware.GetJwt).Methods("GET")
	
	
		//HostRequests
		router.HandleFunc("/event/{id}", middleware.ValidateHostRole(handler.UpdateEvent)).Methods("UPDATE")
		router.HandleFunc("/event/{id}", middleware.ValidateHostRole(handler.DeleteEvent)).Methods("DELETE")	
		router.HandleFunc("/event/invide/{id}", middleware.ValidateHostRole(handler.CreateInventation)).Methods("POST")
		router.HandleFunc("/adduser/{id}/{userid}", middleware.ValidateHostRole(handler.AddParticipants)).Methods("POST")
		
		//Participant Requests
		//Selbst an event teilnehmen
		router.HandleFunc("/event/participate/{id}", middleware.ValidateJWT(handler.Participate)).Methods("POST")
		router.HandleFunc("/event", middleware.ValidateJWT(handler.CreateEvent)).Methods("POST")
		router.HandleFunc("/location",  middleware.ValidateJWT(handler.GetLocations)).Methods("GET")
		router.HandleFunc("/event/{id}",  middleware.ValidateJWT(handler.GetEventById)).Methods("GET")
		router.HandleFunc("/user",middleware.ValidateJWT(handler.GetAllUsers)).Methods("GET")
		router.HandleFunc("/user/{id}",middleware.ValidateJWT(handler.GetUserById)).Methods("GET")
		router.HandleFunc("/user/invetation/{id}",middleware.ValidateJWT(handler.GetInvetations)).Methods("GET")
		router.HandleFunc("/acceptinvetation/{id}",middleware.ValidateJWT(handler.AcceptInvetation)).Methods("POST")
	
		//General User requests	
		router.HandleFunc("/user",handler.CreateUser).Methods("POST")
		router.HandleFunc("/saveEvent/{id}",  handler.SaveEventInCalendar).Methods("POST")
		router.HandleFunc("/event", handler.GetAllEvents).Methods("GET")
		router.HandleFunc("/event/{longitude}/{latitude}",  handler.GetEventsNearby).Methods("GET")
		router.HandleFunc("/eventsofweekend",  handler.GetNextWeekendEvents).Methods("GET")
		router.HandleFunc("/eventsoftoday",  handler.GetEventByDate).Methods("GET")		
		router.HandleFunc("/participants/{id}",handler.GetParticipants).Methods("POST")
	

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