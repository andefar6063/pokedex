package main

import (
	"errors"
	"fmt"
)

func callbackInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("you must provide a specific pokémon name")
	}
	pokemonName := args[0]

	pokemon, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		return errors.New("you haven't caught this pokemon yet")
	}
	fmt.Printf("Name %s\n\n", pokemon.Name)
	fmt.Printf("Heigh %v\n\n", pokemon.Height)
	fmt.Printf("Weight %v\n\n", pokemon.Weight)
	for _, stat := range pokemon.Stats {
		fmt.Printf(" - %s: %v\n\n", stat.Stat.Name, stat.BaseStat)
	}

	for _, typ := range pokemon.Types {
		fmt.Printf(" - %s\n\n", typ.Type.Name)
	}
	return nil
}