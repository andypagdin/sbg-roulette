package main

import (
	"net/http"

	"github.com/google/uuid"
)

type table struct {
	ID      uuid.UUID `json:"id"`
	Players []*player `json:"players"`
}

var tables = make([]*table, 0)

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
