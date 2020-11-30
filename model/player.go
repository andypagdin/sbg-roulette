package model

import (
	"github.com/google/uuid"
)

type Player struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Balance float64   `json:"balance"`
}

var Players = make([]*Player, 0)

func (p *Player) AddBalance(a float64) {
	p.Balance += a
}

func (p *Player) DeductBalance(a float64) {
	p.Balance -= a
}

func (p *Player) AddPlayer() {
	p.ID = uuid.New()
	p.Balance = 100
	Players = append(Players, p)
}

func GetPlayer(id string) (*Player, string) {
	for _, n := range Players {
		if n.ID.String() == id {
			return n, ""
		}
	}
	return nil, "Player not found"
}
