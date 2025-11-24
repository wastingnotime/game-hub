package handlers

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/wastingnotime/game-hub/domain"

	"net/http"
	"time"
)

func StartDuelHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		PlayerID string `json:"player_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	challenge := domain.Challenge{
		ID:        uuid.New().String(),
		PlayerID:  req.PlayerID,
		Status:    "waiting",
		CreatedAt: time.Now().UTC(),
	}

	// TODO: Save challenge (in-memory, Redis, or DB)
	// For now, just return it
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(challenge)
}
