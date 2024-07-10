// A simple API for legendary animals.

package main

import (
	"context"
	"log"
	"net/http"
	"encoding/json"
	"strconv"

	"github.com/gorilla/mux"
)

type Legend struct {
	ID			string `json:"id"`
	Name		string `json:"name"`
	Type    string `json:"type"`
	Origin 	string `json:"origin"`
}

var legends []Legend

func getLegends(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(legends)
}

func createLegend(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newLegend Legend
	json.NewDecoder(r.Body).Decode(&newLegend)
	newLegend.ID = strconv.Itoa(len(legends) +1)
	legends = append(legends, newLegend)
	json.NewEncoder(w).Encode(newLegend)
}

func main() {
	if err := run(context.Background()); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context) error {
	router:=mux.NewRouter()
	router.HandleFunc("/legend", getLegends).Methods("GET")
	router.HandleFunc("/legend", createLegend).Methods("POST")

	return http.ListenAndServe(":5000", router)
}
