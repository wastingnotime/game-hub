package domain

import (
	"time"

	"github.com/google/uuid"
)

type Duel struct {
	ID        string    `json:"id"`
	Players   []string  `json:"players"`
	State     string    `json:"state"` // starting, active, finished
	CreatedAt time.Time `json:"created_at"`
}

func NewDuel(p1, p2 string) Duel {
	return Duel{
		ID:        uuid.New().String(),
		Players:   []string{p1, p2},
		State:     "starting",
		CreatedAt: time.Now().UTC(),
	}
}
