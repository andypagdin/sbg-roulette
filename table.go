package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type table struct {
	ID      uuid.UUID `json:"id"`
	Players []*player `json:"players"`
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
	respondWithJSON(w, http.StatusOK, tables)
}

func tablesHandlerPost(w http.ResponseWriter, r *http.Request) {
	table := new(table)
	table.ID = uuid.New()
	table.Players = make([]*player, 0)

	tables = append(tables, table)
	respondWithJSON(w, http.StatusOK, table)
}

func tablesHandlePlayerPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	table, err := getTable(vars["table-id"])
	if err != "" {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	err = isPlayerAtTable(table, vars["player-id"])
	if err != "" {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	player, err := getPlayer(vars["player-id"])

	if err != "" {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	table.Players = append(table.Players, player)
	respondWithJSON(w, http.StatusOK, table)
}
