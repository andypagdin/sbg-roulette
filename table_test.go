package main

import (
	"encoding/json"
	"net/http"
	"testing"
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
