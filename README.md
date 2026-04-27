# game-hub

Game Hub — a platform designed to host multiple simple games under one structure

## Status: Dormant / Paused

GameHub explored state, matchmaking, timing, and interaction patterns through the Trivia Duel concept.
The project is paused because its feedback loop is too slow and open-ended for the current phase.
It is preserved as a learning artifact, not deleted.

## What this project explored

- matchmaking
- duel lifecycle
- state coordination
- interaction timing
- simple game backend architecture

## Why it is paused

- no bounded finish line
- slow feedback loop
- not aligned with current execution energy
- preserved for future reference

## about the project 

It a slow walkthrough to build a game platform that can be found at:
https://wastingnotime.org/sagas/game-hub/


## get starting

to try the app you need to start the api
```bash
cd apps/api

go mod download

go run main.go
```

so in another terminal window you can type:
```bash
curl -X POST http://localhost:8080/duel/start \
  -H "Content-Type: application/json" \
  -d '{"player_id": "P001"}'
```
