package handler

import (
	"encoding/json"
	"net/http"

	"github.com/andypagdin/sbg-roulette/model"
	"github.com/google/uuid"
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

	// todo: move into player model
	player := new(model.Player)
	player.ID = uuid.New()
	player.Name = p.Name
	player.Balance = 100

	// todo: move into player model
	model.Players = append(model.Players, player)
	RespondJSON(w, http.StatusOK, player)
}
