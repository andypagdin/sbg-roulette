package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/andypagdin/sbg-roulette/model"
)

func TestPlayersHandlerGet(t *testing.T) {
	clearPlayers()

	req, _ := http.NewRequest("GET", "/v1/players", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	body := response.Body.String()

	if body != "[]" {
		t.Errorf("Expected an empty array. Got '%s'", body)
	}
}

func TestPlayersHandlerPost(t *testing.T) {
	clearPlayers()

	var jsonStr = []byte(`{"Name": "Foo"}`)
	req, _ := http.NewRequest("POST", "/v1/players", bytes.NewBuffer(jsonStr))

	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	var m model.Player
	_ = json.Unmarshal(response.Body.Bytes(), &m)

	if m.ID != model.Players[0].ID {
		t.Errorf("Expected player ID to be '%s'. Got '%s'", m.ID, model.Players[0].ID)
	}

	if m.Name != "Foo" {
		t.Errorf("Expected player Name to be 'Foo'. Got '%s'", m.Name)
	}
}
