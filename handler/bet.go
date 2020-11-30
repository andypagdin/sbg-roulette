package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/andypagdin/sbg-roulette/model"
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

func tablesBetHandlerPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var b model.Bet

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&b)

	if err != nil {
		RespondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	table, err2 := model.GetTable(vars["table-id"])
	if err2 != "" {
		RespondError(w, http.StatusBadRequest, err2)
		return
	}
	if !table.OpenForBets {
		RespondError(w, http.StatusBadRequest, "Bets are closed wait for next round")
		return
	}

	err3 := model.IsPlayerAtTable(table, vars["player-id"])
	if err3 == "" {
		RespondError(w, http.StatusBadRequest, "Player must be added to the table before placing a bet")
		return
	}

	player, err4 := model.GetPlayer(vars["player-id"])
	if err4 != "" {
		RespondError(w, http.StatusBadRequest, err4)
		return
	}

	uuid, _ := uuid.Parse(vars["player-id"])

	// todo: move into bets model
	bet := new(model.Bet)
	bet.PlayerID = uuid
	bet.Type = b.Type
	bet.Value = b.Value
	bet.Amount = b.Amount

	// todo: move into table model
	table.Bets = append(table.Bets, bet)

	// todo: move into player model
	player.Balance -= b.Amount

	RespondJSON(w, http.StatusOK, bet)
}

func tablesBetSettleHandlerPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	table, err1 := model.GetTable(vars["table-id"])
	if err1 != "" {
		RespondError(w, http.StatusBadRequest, err1)
		return
	}

	outcome, err2 := strconv.Atoi(vars["outcome"])
	if err2 != nil {
		RespondError(w, http.StatusBadRequest, "Invalid outcome paramater")
		return
	}

	for _, b := range table.Bets {
		result := model.GetBetOutcome(b, outcome)
		model.SettleBet(result, b.PlayerID.String())
	}

	// todo: move into bets model
	table.Bets = make([]*model.Bet, 0)
	table.OpenForBets = true
}
