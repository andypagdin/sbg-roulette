package model

import (
	"strconv"

	"github.com/google/uuid"
)

type Bet struct {
	PlayerID uuid.UUID `json:"playerId"`
	Type     string    `json:"type"`
	Value    string    `json:"value"`
	Amount   float64   `json:"amount"`
}

func GetBetOutcome(b *Bet, o int) float64 {
	switch b.Type {
	case "straight":
		v, _ := strconv.Atoi(b.Value)
		return getStraightBetOutcome(v, o, b.Amount)
	case "colour":
		return getColourBetOutcome(b.Value, o, b.Amount)
	case "oddEven":
		return getOddEvenBetOutcome(b.Value, o, b.Amount)
	case "highLow":
		return getHighLowBetOutcome(b.Value, o, b.Amount)
	case "column":
		return getColumnBetOutcome(b.Value, o, b.Amount)
	}

	return 0
}

func SettleBet(r float64, id string) {
	player, _ := GetPlayer(id)
	player.AddBalance(r)
}

func getStraightBetOutcome(v int, o int, a float64) float64 {
	if v == o {
		return a * 35
	}
	return 0
}

func getColourBetOutcome(v string, o int, a float64) float64 {
	if v == "red" && o%2 == 0 || v == "black" && o%2 != 0 {
		return a * 2
	}
	return 0
}

func getOddEvenBetOutcome(v string, o int, a float64) float64 {
	if v == "even" && o%2 == 0 || v == "odd" && o%2 != 0 {
		return a * 2
	}
	return 0
}

func getHighLowBetOutcome(v string, o int, a float64) float64 {
	if v == "low" && o >= 1 && o <= 18 || v == "high" && o >= 19 && o <= 36 {
		return a * 2
	}
	return 0
}

func getColumnBetOutcome(v string, o int, a float64) float64 {
	if v == "1st12" && o >= 1 && o <= 12 ||
		v == "2nd12" && o >= 13 && o <= 22 ||
		v == "3rd12" && o >= 25 && o <= 34 {
		return a * 2
	}
	return 0
}
