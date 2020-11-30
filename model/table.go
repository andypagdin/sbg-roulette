package model

import "github.com/google/uuid"

type Table struct {
	ID          uuid.UUID `json:"id"`
	Players     []*Player `json:"players"`
	Bets        []*Bet    `json:"bets"`
	OpenForBets bool      `json:"openForBets"`
}

var Tables = make([]*Table, 0)

func GetTable(id string) (*Table, string) {
	for _, n := range Tables {
		if n.ID.String() == id {
			return n, ""
		}
	}
	return nil, "Table not found"
}

func IsPlayerAtTable(table *Table, playerID string) string {
	for _, n := range table.Players {
		if n.ID.String() == playerID {
			return "Player is already at this table"
		}
	}
	return ""
}
