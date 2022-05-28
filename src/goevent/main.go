package main

import (
	"net/http"

	"github.com/RomanoLu/events-to-go/src/goevent/db"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Println("Starting My-Aktion API server")
	router := mux.NewRouter()
	//router.HandleFunc("/campaign", ).Methods("POST")
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
