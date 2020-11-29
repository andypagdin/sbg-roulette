package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", router()))
}

func router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/v1/tables", tablesHandlerGet).Methods("GET")
	r.HandleFunc("/v1/tables", tablesHandlerPost).Methods("POST")
	r.HandleFunc("/v1/tables/{table-id}/spin", tablesSpinHandlerGet).Methods("GET")
	r.HandleFunc("/v1/tables/{table-id}/bet/{player-id}", tablesBetHandlerPost).Methods("POST")
	r.HandleFunc("/v1/tables/{table-id}/players/{player-id}", tablesPlayerHandlerPost).Methods("POST")
	r.HandleFunc("/v1/players", playersHandlerGet).Methods("GET")
	r.HandleFunc("/v1/players", playersHandlerPost).Methods("POST")
	return r
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)

	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	_, err = w.Write(response)

	if err != nil {
		log.Fatal(err)
	}
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}
