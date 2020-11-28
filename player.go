package main

import "github.com/google/uuid"

type player struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
