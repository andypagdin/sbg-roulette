package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// Assumption: All incoming bets will have a correctly set type
type bet struct {
	PlayerID uuid.UUID `json:"playerId"`
	Type     string    `json:"type"`
	Value    string    `json:"value"`
	Amount   float64   `json:"amount"`
}

func tablesBetHandlerPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var b bet

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&b)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	table, err2 := getTable(vars["table-id"])
	if err2 != "" {
		respondWithError(w, http.StatusBadRequest, err2)
		return
	}
	if !table.OpenForBets {
		respondWithError(w, http.StatusBadRequest, "Bets are closed wait for next round")
		return
	}

	err3 := isPlayerAtTable(table, vars["player-id"])
	if err3 == "" {
		respondWithError(w, http.StatusBadRequest, "Player must be added to the table before placing a bet")
		return
	}

	uuid, err4 := uuid.Parse(vars["player-id"])
	if err4 != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid player ID")
		return
	}

	bet := new(bet)
	bet.PlayerID = uuid
	bet.Type = b.Type
	bet.Value = b.Value
	bet.Amount = b.Amount

	table.Bets = append(table.Bets, bet)
	respondWithJSON(w, http.StatusOK, bet)
}
