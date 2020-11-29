package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestTablesBetHandlerPost(t *testing.T) {
	clearTables()
	clearPlayers()

	tbl := addTable()
	plr := addPlayer()
	addPlayerToTable(plr, tbl)

	var jsonStr = []byte(`{"type": "straight", "value": "10", "amount": 50}`)
	req, _ := http.NewRequest("POST", "/v1/tables/"+tbl.ID.String()+"/bet/"+plr.ID.String(), bytes.NewBuffer(jsonStr))

	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	var b bet
	_ = json.Unmarshal(response.Body.Bytes(), &b)

	if b.PlayerID != plr.ID {
		t.Errorf("Expected bet PlayerID to be '%s'. Got '%s'", plr.ID, b.PlayerID)
	}

	if b.Amount != 50 {
		t.Errorf("Expected bet Amount to be '50'. Got '%f'", b.Amount)
	}

	if b.Type != "straight" {
		t.Errorf("Expected bet Type to be 'straight'. Got '%s'", b.Type)
	}

	if b.Value != "10" {
		t.Errorf("Expected bet Value to be '10'. Got '%s'", b.Value)
	}
}

func TestTablesBetSettleHandlerPost(t *testing.T) {
	clearTables()
	clearPlayers()

	// Use a new player for each bet type to keep assertions simple
	tbl := addTable()
	plr1 := addPlayer()
	plr2 := addPlayer()
	plr3 := addPlayer()
	plr4 := addPlayer()
	plr5 := addPlayer()
	plr6 := addPlayer()

	addPlayerToTable(plr1, tbl)
	addPlayerToTable(plr2, tbl)
	addPlayerToTable(plr3, tbl)
	addPlayerToTable(plr4, tbl)
	addPlayerToTable(plr5, tbl)
	addPlayerToTable(plr6, tbl)

	addBetToTable(plr1, tbl, "straight", "10", 100)
	addBetToTable(plr2, tbl, "straight", "5", 100)
	addBetToTable(plr3, tbl, "colour", "red", 100)
	addBetToTable(plr4, tbl, "colour", "black", 100)
	addBetToTable(plr5, tbl, "oddEven", "even", 100)
	addBetToTable(plr6, tbl, "oddEven", "odd", 100)

	req, _ := http.NewRequest("POST", "/v1/tables/"+tbl.ID.String()+"/bet/settle/10", nil)

	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	if plr1.Balance != 3500 {
		t.Errorf("Expected plr1 win balance '3500'. Got %.2f", plr1.Balance)
	}

	if plr2.Balance != 0 {
		t.Errorf("Expected plr2 lose balance '0'. Got %.2f", plr2.Balance)
	}

	if plr3.Balance != 200 {
		t.Errorf("Expected plr3 win balance '200'. Got %.2f", plr3.Balance)
	}

	if plr4.Balance != 0 {
		t.Errorf("Expected plr4 lose balance '0'. Got %.2f", plr4.Balance)
	}

	if plr5.Balance != 200 {
		t.Errorf("Expected plr5 win balance '200'. Got %.2f", plr5.Balance)
	}

	if plr6.Balance != 0 {
		t.Errorf("Expected plr6 lose balance '0'. Got %.2f", plr6.Balance)
	}

	if len(tbl.Bets) != 0 {
		t.Errorf("Expected table bets to be settled '0'. Got %d", len(tbl.Bets))
	}

	if tbl.OpenForBets != true {
		t.Errorf("Expected table OpenForBets 'false'. Got %t", tbl.OpenForBets)
	}
}
