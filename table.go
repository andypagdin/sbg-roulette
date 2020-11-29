package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/andypagdin/sbg-roulette/respond"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type table struct {
	ID          uuid.UUID `json:"id"`
	Players     []*player `json:"players"`
	Bets        []*bet    `json:"bets"`
	OpenForBets bool      `json:"openForBets"`
}

var tables = make([]*table, 0)

func getTable(id string) (*table, string) {
	for _, n := range tables {
		if n.ID.String() == id {
			return n, ""
		}
	}
	return nil, "Table not found"
}

func isPlayerAtTable(table *table, playerID string) string {
	for _, n := range table.Players {
		if n.ID.String() == playerID {
			return "Player is already at this table"
		}
	}
	return ""
}

func tablesHandlerGet(w http.ResponseWriter, r *http.Request) {
	respond.JSON(w, http.StatusOK, tables)
}

func tablesHandlerPost(w http.ResponseWriter, r *http.Request) {
	table := new(table)
	table.ID = uuid.New()
	table.Players = make([]*player, 0)
	table.Bets = make([]*bet, 0)
	table.OpenForBets = true

	tables = append(tables, table)
	respond.JSON(w, http.StatusOK, table)
}

func tablesPlayerHandlerPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	table, err := getTable(vars["table-id"])
	if err != "" {
		respond.Error(w, http.StatusBadRequest, err)
		return
	}

	err = isPlayerAtTable(table, vars["player-id"])
	if err != "" {
		respond.Error(w, http.StatusBadRequest, err)
		return
	}

	player, err := getPlayer(vars["player-id"])
	if err != "" {
		respond.Error(w, http.StatusBadRequest, err)
		return
	}

	table.Players = append(table.Players, player)
	respond.JSON(w, http.StatusOK, table)
}

func tablesSpinHandlerGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	table, err := getTable(vars["table-id"])
	if err != "" {
		respond.Error(w, http.StatusBadRequest, err)
		return
	}
	if !table.OpenForBets {
		respond.Error(w, http.StatusBadRequest, "Settle outstanding bets before spinning")
		return
	}

	table.OpenForBets = false

	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 36
	outcome := rand.Intn(max-min+1) + min

	respond.JSON(w, http.StatusOK, outcome)
}
