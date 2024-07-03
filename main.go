package main

import "github.com/andefar6063/pokedex/internal/pokeapi"

type config struct {
	pokeapiClient pokeapi.Client
	nextLocationAreaURL *string
	prevLoactaionAreaURL *string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(),
	}

	startRepl(&cfg)
}