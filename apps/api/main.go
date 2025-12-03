package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/wastingnotime/game-hub/handlers"
	"github.com/wastingnotime/game-hub/services"
)

func main() {
	fmt.Println("🕹️ Game Hub starting...")
	services.StartMatchmaker()

	//POST /duel/start
	http.HandleFunc("/duel/start", handlers.StartDuelHandler)

	s := &http.Server{
		Addr: ":8080",
	}

	log.Fatal(s.ListenAndServe())
}
