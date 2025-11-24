## desired folder strcuture for game-hub monorepo

This folder structure is just a draft roadmap. I won’t create everything upfront — the idea is to let the project unfold as it grows.

Some choices here are pre-defined guesses, not final decisions. Trivia Duel is the first game in GameHub, so it’s reasonable to expect its logic to be split from the core hub over time.

But I’m not trying to play too many moves ahead: this is just one probable direction, not a commitment. At any point we might face a plot twist and take a different path.

game-hub/
  README.md                      # narrative + quickstart
  apps/
    api/                         # game-hub backend
    web/                         # frontend
    trivia-duel-service/         # trivia domain logic
  libs/
    core/                        # shared domain models, utils
    game-sdk/                    # helpers to build new games
  infra/
    docker/
    k8s/
  docs/
    01-overview.md
    02-architecture.md
    10-trivia-duel.md
