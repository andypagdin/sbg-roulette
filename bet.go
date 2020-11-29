package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

/**
	Assumption: All incoming bets will have a correctly set type.
	This could be ensured by exposing a 'get board' endpoint which would
	return all valid board tiles with corresponding bet types and values for
	the implementer generate the front end with, ensuring the correct
	bet type is sent on front end interaction.
**/
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

	player, err4 := getPlayer(vars["player-id"])
	if err4 != "" {
		respondWithError(w, http.StatusBadRequest, err4)
		return
	}

	uuid, _ := uuid.Parse(vars["player-id"])

	bet := new(bet)
	bet.PlayerID = uuid
	bet.Type = b.Type
	bet.Value = b.Value
	bet.Amount = b.Amount

	table.Bets = append(table.Bets, bet)

	player.Balance -= b.Amount

	respondWithJSON(w, http.StatusOK, bet)
}

func tablesBetSettleHandlerPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	table, err1 := getTable(vars["table-id"])
	if err1 != "" {
		respondWithError(w, http.StatusBadRequest, err1)
		return
	}

	outcome, err2 := strconv.Atoi(vars["outcome"])
	if err2 != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid outcome paramater")
		return
	}

	for _, b := range table.Bets {
		result := getBetOutcome(b, outcome)
		settleBet(result, b.PlayerID.String())
	}

	table.Bets = make([]*bet, 0)
	table.OpenForBets = true
}

func getBetOutcome(b *bet, o int) float64 {
	switch b.Type {
	case "straight":
		v, _ := strconv.Atoi(b.Value)
		return getStraightBetOutcome(v, o, b.Amount)
	case "colour":
		return getColourBetOutcome(b.Value, o, b.Amount)
	case "oddEven":
		return getOddEvenBetOutcome(b.Value, o, b.Amount)
	case "highLow":
		return getHighLowBetOutcome(b.Value, o, b.Amount)
	}

	return 0
}

func getStraightBetOutcome(v int, o int, a float64) float64 {
	if v == o {
		return a * 35
	}
	return 0
}

func getColourBetOutcome(v string, o int, a float64) float64 {
	if v == "red" && o%2 == 0 || v == "black" && o%2 != 0 {
		return a * 2
	}
	return 0
}

func getOddEvenBetOutcome(v string, o int, a float64) float64 {
	if v == "even" && o%2 == 0 || v == "odd" && o%2 != 0 {
		return a * 2
	}
	return 0
}

func getHighLowBetOutcome(v string, o int, a float64) float64 {
	if v == "low" && o >= 1 && o <= 18 || v == "high" && o >= 19 && o <= 36 {
		return a * 2
	}
	return 0
}

func settleBet(r float64, id string) {
	player, _ := getPlayer(id)
	player.Balance += r
}
