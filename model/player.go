package model

import "github.com/google/uuid"

type Player struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Balance float64   `json:"balance"`
}

var Players = make([]*Player, 0)

func GetPlayer(id string) (*Player, string) {
	for _, n := range Players {
		if n.ID.String() == id {
			return n, ""
		}
	}
	return nil, "Player not found"
}
