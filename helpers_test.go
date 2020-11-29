package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
)

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	router().ServeHTTP(rr, req)
	return rr
}

func clearTables() {
	tables = make([]*table, 0)
}

func clearPlayers() {
	players = make([]*player, 0)
}

func addTable() *table {
	table := new(table)
	table.ID = uuid.New()
	table.Players = make([]*player, 0)
	table.Bets = make([]*bet, 0)
	table.OpenForBets = true
	tables = append(tables, table)
	return table
}

func addPlayer() *player {
	player := new(player)
	player.ID = uuid.New()
	player.Name = "Foo"
	player.Balance = 100
	players = append(players, player)
	return player
}

func addPlayerToTable(p *player, t *table) {
	t.Players = append(t.Players, p)
}

func addBetToTable(p *player, t *table, bType string, bValue string, bAmount float64) {
	bet := new(bet)
	bet.PlayerID = p.ID
	bet.Type = bType
	bet.Value = bValue
	bet.Amount = bAmount
	p.Balance -= bAmount
	t.Bets = append(t.Bets, bet)
}
