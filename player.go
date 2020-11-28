package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type player struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

var players = make([]*player, 0)

func getPlayer(id string) (*player, string) {
	for _, n := range players {
		if n.ID.String() == id {
			return n, ""
		}
	}
	return nil, "Player not found"
}

func playersHandlerGet(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, players)
}

func playersHandlerPost(w http.ResponseWriter, r *http.Request) {
	var p player
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	plr := new(player)
	plr.ID = uuid.New()
	plr.Name = p.Name

	players = append(players, plr)
	respondWithJSON(w, http.StatusOK, plr)
}
