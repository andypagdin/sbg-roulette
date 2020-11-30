package handler

import (
	"encoding/json"
	"net/http"

	"github.com/andypagdin/sbg-roulette/model"
)

func playersHandlerGet(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, model.Players)
}

func playersHandlerPost(w http.ResponseWriter, r *http.Request) {
	var p model.Player
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)

	if err != nil {
		RespondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	p.AddPlayer()
	RespondJSON(w, http.StatusOK, p)
}
