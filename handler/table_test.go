package handler

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/andypagdin/sbg-roulette/model"
)

func TestTablesHandlerGet(t *testing.T) {
	clearTables()

	req, _ := http.NewRequest("GET", "/v1/tables", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	body := response.Body.String()

	if body != "[]" {
		t.Errorf("Expected an empty array. Got '%s'", body)
	}
}

func TestTablesHandlerPost(t *testing.T) {
	clearTables()

	req, _ := http.NewRequest("POST", "/v1/tables", nil)

	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	var m model.Table
	_ = json.Unmarshal(response.Body.Bytes(), &m)

	if m.ID != model.Tables[0].ID {
		t.Errorf("Expected table ID to be '%s'. Got '%s'", m.ID, model.Tables[0].ID)
	}
}

func TestTablesPlayerHandlerPost(t *testing.T) {
	clearTables()
	clearPlayers()

	tbl := addTable()
	plr := addPlayer()

	req, _ := http.NewRequest("POST", "/v1/tables/"+tbl.ID.String()+"/players/"+plr.ID.String(), nil)

	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	var m model.Table
	_ = json.Unmarshal(response.Body.Bytes(), &m)

	if m.Players[0].ID != plr.ID {
		t.Errorf("Expected table player ID '%s'. Got '%s'", plr.ID, m.Players[0].ID)
	}

	if m.Players[0].Name != "Foo" {
		t.Errorf("Expected table player Name 'Foo'. Got '%s'", m.Players[0].Name)
	}
}

func TestTablesSpinHandlerGet(t *testing.T) {
	clearTables()
	clearPlayers()

	tbl := addTable()

	req, _ := http.NewRequest("GET", "/v1/tables/"+tbl.ID.String()+"/spin", nil)

	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	var o int
	_ = json.Unmarshal(response.Body.Bytes(), &o)

	if o < 0 || o > 36 {
		t.Errorf("Expected outcome to within game rules '0-36'. Got '%d'", o)
	}

	if tbl.OpenForBets {
		t.Errorf("Expected table OpenForBets to be 'false'. Got '%t'", tbl.OpenForBets)
	}
}
