package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/OneTaxApi/location-taxrate-api/api/datastore"
	"log"
	"net/http"
	"time"
)

var (
	taxrate datastore.location
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func init() {
	defer timeTrack(time.Now(), "file load")
	taxrates = &datastore.location{}
	taxrates.Initialize()
}

func main() {
	r := mux.NewRouter()
	log.Println("taxrate api")
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "api v1")
	})
	api.HandleFunc("/taxrate/location/{location}", searchByLocation).Methods(http.MethodGet)
	log.Fatalln(http.ListenAndServe(":8080", r))
}
