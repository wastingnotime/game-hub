package main

import (
	"log"
	"net/http"

	"github.com/wastingnotime/game-hub/handlers"
)

func main() {

	//POST /duel/start
	http.HandleFunc("/duel/start", handlers.StartDuelHandler)

	s := &http.Server{
		Addr: ":8080",
	}

	log.Fatal(s.ListenAndServe())
}
