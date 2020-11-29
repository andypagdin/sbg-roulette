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
