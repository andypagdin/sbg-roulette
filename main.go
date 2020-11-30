package main

import (
	"log"
	"net/http"

	"github.com/andypagdin/sbg-roulette/handler"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	handler.RegisterRouteHandlers(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}
