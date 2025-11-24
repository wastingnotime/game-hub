# game-hub

Game Hub — a platform designed to host multiple simple games under one structure

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