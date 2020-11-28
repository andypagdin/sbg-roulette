package main

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/google/uuid"
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

	var m table
	_ = json.Unmarshal(response.Body.Bytes(), &m)

	if m.ID != tables[0].ID {
		t.Errorf("Expected table ID to be '%s'. Got '%s'", m.ID, tables[0].ID)
	}
}

func TestTablesHandlerPlayerPost(t *testing.T) {
	clearTables()
	clearPlayers()

	tbl := addTable()
	plr := addPlayer()

	req, _ := http.NewRequest("POST", "/v1/tables/"+tbl.ID.String()+"/players/"+plr.ID.String(), nil)

	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	var m table
	_ = json.Unmarshal(response.Body.Bytes(), &m)

	if m.Players[0].ID != plr.ID {
		t.Errorf("Expected table player ID '%s'. Got '%s'", plr.ID, m.Players[0].ID)
	}

	if m.Players[0].Name != "Foo" {
		t.Errorf("Expected table player Name 'Foo'. Got '%s'", m.Players[0].Name)
	}
}

func TestTableFound(t *testing.T) {
	clearTables()

	tbl := addTable()
	foundTbl, err := getTable(tbl.ID.String())

	if err != "" {
		t.Errorf("Expected table to be found and err to be ''. Got '%s'", err)
	}

	if foundTbl.ID != tbl.ID {
		t.Errorf("Expected table ID '%s'. Got '%s'", tbl.ID, foundTbl.ID)
	}
}

func TestGetTable(t *testing.T) {
	clearTables()

	tbl := addTable()
	foundTbl, foundTblErr := getTable(tbl.ID.String())
	notFoundTbl, notFoundTblErr := getTable("123-456-789")

	if foundTblErr != "" {
		t.Errorf("Expected table to be found and err to be ''. Got '%s'", foundTblErr)
	}

	if foundTbl.ID != tbl.ID {
		t.Errorf("Expected table ID '%s'. Got '%s'", tbl.ID, foundTbl.ID)
	}

	if notFoundTblErr != "Table not found" {
		t.Errorf("Expected table error 'Table not found'. Got '%s'", notFoundTblErr)
	}

	if notFoundTbl != nil {
		t.Errorf("Expected table to be 'Nil'. Got '%v'", notFoundTbl)
	}
}

func TestIsPlayerAtTable(t *testing.T) {
	clearTables()
	clearPlayers()

	tbl := addTable()
	plr := addPlayer()

	tbl.Players = append(tbl.Players, plr)

	err1 := isPlayerAtTable(tbl, plr.ID.String())
	err2 := isPlayerAtTable(tbl, uuid.New().String())

	if err1 != "Player is already at this table" {
		t.Errorf("Expected 'Player is already at this table'. Got '%s'", err1)
	}

	if err2 != "" {
		t.Errorf("Expected player to not be at table ''. Got '%s'", err2)
	}
}
