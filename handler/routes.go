package handler

import "github.com/gorilla/mux"

func RegisterRouteHandlers(r *mux.Router) {
	r.HandleFunc("/v1/tables", tablesHandlerGet).Methods("GET")
	r.HandleFunc("/v1/tables", tablesHandlerPost).Methods("POST")
	r.HandleFunc("/v1/tables/{table-id}/spin", tablesSpinHandlerGet).Methods("GET")
	r.HandleFunc("/v1/tables/{table-id}/bet/{player-id}", tablesBetHandlerPost).Methods("POST")
	r.HandleFunc("/v1/tables/{table-id}/bet/settle/{outcome}", tablesBetSettleHandlerPost).Methods("POST")
	r.HandleFunc("/v1/tables/{table-id}/players/{player-id}", tablesPlayerHandlerPost).Methods("POST")
	r.HandleFunc("/v1/players", playersHandlerGet).Methods("GET")
	r.HandleFunc("/v1/players", playersHandlerPost).Methods("POST")
}
