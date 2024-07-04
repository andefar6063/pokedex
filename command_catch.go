package main

import (
	"fmt"
	"math/rand"
)

func callbackPokemon(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("you must provide a specific pokémon name")
	}
	pokemonName := args[0]

	pokemonCatch, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	const threshold = 50
	randNum := rand.Intn(pokemonCatch.BaseExperience)
	if randNum > threshold {
		return fmt.Errorf("failed to catch %s", pokemonName)
	}

	cfg.caughtPokemon[pokemonName] = pokemonCatch
	fmt.Printf("Caught %s\n", pokemonName)

	return nil
}