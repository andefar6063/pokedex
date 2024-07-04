package main

import (
	"fmt"
)

func callbackInspectPokedex(cfg *config, args ...string) error {
	if len(cfg.caughtPokemon) == 0 {
		return fmt.Errorf("you do not have any pokemons at the moment")
	}
	for _, i := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n\n", i.Name)
	}
	return nil
}