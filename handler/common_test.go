package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andypagdin/sbg-roulette/model"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	RegisterRouteHandlers(r)
	r.ServeHTTP(rr, req)
	return rr
}

/**
	Helper functions to create common data structures for use in tests
**/

func clearTables() {
	model.Tables = make([]*model.Table, 0)
}

func clearPlayers() {
	model.Players = make([]*model.Player, 0)
}

func addTable() *model.Table {
	table := new(model.Table)
	table.ID = uuid.New()
	table.Players = make([]*model.Player, 0)
	table.Bets = make([]*model.Bet, 0)
	table.OpenForBets = true
	model.Tables = append(model.Tables, table)
	return table
}

func addPlayer() *model.Player {
	player := new(model.Player)
	player.ID = uuid.New()
	player.Name = "Foo"
	player.Balance = 100
	model.Players = append(model.Players, player)
	return player
}

func addPlayerToTable(p *model.Player, t *model.Table) {
	t.Players = append(t.Players, p)
}

func addBetToTable(p *model.Player, t *model.Table, bType string, bValue string, bAmount float64) {
	bet := new(model.Bet)
	bet.PlayerID = p.ID
	bet.Type = bType
	bet.Value = bValue
	bet.Amount = bAmount
	p.Balance -= bAmount
	t.Bets = append(t.Bets, bet)
}
