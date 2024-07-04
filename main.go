package main

import (
	"time"

	"github.com/andefar6063/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	nextLocationAreaURL *string
	prevLoactaionAreaURL *string
	caughtPokemon map[string]pokeapi.Pokemon
}

func main() {
	cacheInterval := 5 * time.Minute

	cfg := config{
		pokeapiClient: pokeapi.NewClient(cacheInterval),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}

	startRepl(&cfg)
}