package main

import (
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
	r.HandleFunc("/v1/tables/{table-id}/bet/settle/{outcome}", tablesBetSettleHandlerPost).Methods("POST")
	r.HandleFunc("/v1/tables/{table-id}/players/{player-id}", tablesPlayerHandlerPost).Methods("POST")
	r.HandleFunc("/v1/players", playersHandlerGet).Methods("GET")
	r.HandleFunc("/v1/players", playersHandlerPost).Methods("POST")
	return r
}
