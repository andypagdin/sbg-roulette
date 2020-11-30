package handler

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/andypagdin/sbg-roulette/model"
	"github.com/gorilla/mux"
)

func tablesHandlerGet(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, model.Tables)
}

func tablesHandlerPost(w http.ResponseWriter, r *http.Request) {
	var t model.Table
	t.AddTable()
	RespondJSON(w, http.StatusOK, t)
}

func tablesPlayerHandlerPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	table, err := model.GetTable(vars["table-id"])
	if err != "" {
		RespondError(w, http.StatusBadRequest, err)
		return
	}

	err = model.IsPlayerAtTable(table, vars["player-id"])
	if err != "" {
		RespondError(w, http.StatusBadRequest, err)
		return
	}

	player, err := model.GetPlayer(vars["player-id"])
	if err != "" {
		RespondError(w, http.StatusBadRequest, err)
		return
	}

	table.AddPlayerToTable(player)
	RespondJSON(w, http.StatusOK, table)
}

func tablesSpinHandlerGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	table, err := model.GetTable(vars["table-id"])
	if err != "" {
		RespondError(w, http.StatusBadRequest, err)
		return
	}
	if !table.OpenForBets {
		RespondError(w, http.StatusBadRequest, "Settle outstanding bets before spinning")
		return
	}

	table.SetOpenForBets(false)

	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 36
	outcome := rand.Intn(max-min+1) + min

	RespondJSON(w, http.StatusOK, outcome)
}
