package services

import (
	"fmt"
	"sync"
	"time"

	"github.com/wastingnotime/game-hub/domain"
	"github.com/wastingnotime/game-hub/memory"
)

var mu sync.Mutex

func StartMatchmaker() {
	go func() {
		for {
			time.Sleep(2 * time.Second) // small heartbeat
			matchPlayers()
		}
	}()
}

func matchPlayers() {
	mu.Lock()
	defer mu.Unlock()

	for len(memory.WaitingChallenges) >= 2 {
		p1 := memory.WaitingChallenges[0]
		p2 := memory.WaitingChallenges[1]
		memory.WaitingChallenges = memory.WaitingChallenges[2:]

		duel := domain.NewDuel(p1.PlayerID, p2.PlayerID)
		fmt.Printf("✨ New duel started! %s vs %s → ID: %s\n", p1.PlayerID, p2.PlayerID, duel.ID)
	}
}
