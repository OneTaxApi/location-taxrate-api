package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/OneTaxApi/location-taxrate-api/location-taxrate-api/api/loader"
	"net/http"
	"strconv"
)

func searchByLocation(w http.ResponseWriter, r *http.Request) {
	queries := mux.Vars(r)

	w.Header().Set("Content-Type", "application/json")
	if val, ok := queries["location"]; ok {
		data := taxrates.SearchLocation(location)
		if data != nil {
			b, err := json.Marshal(data)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{"error": "error marshalling data"}`))
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(b)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"error": "not found"}`))
}

