package main

import (
	"time"

	"github.com/andefar6063/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	nextLocationAreaURL *string
	prevLoactaionAreaURL *string
}

func main() {
	cacheInterval := 5 * time.Minute

	cfg := config{
		pokeapiClient: pokeapi.NewClient(cacheInterval),
	}

	startRepl(&cfg)
}