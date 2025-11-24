package domain

import "time"

type Challenge struct {
	ID        string    `json:"id"`
	PlayerID  string    `json:"player_id"`
	Status    string    `json:"status"` // waiting, matched, cancelled
	CreatedAt time.Time `json:"created_at"`
}
