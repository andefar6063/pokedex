package main

import (
	"errors"
	"fmt"
	"log"
)

func callbackMap(cfg *config, args ...string) error {

	res, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Location areas:")
	for _, area := range res.Results {
		fmt.Printf("- %s\n", area.Name)
	}
	cfg.nextLocationAreaURL = res.Next
	cfg.prevLoactaionAreaURL = res.Previous
	return nil
}

func callbackMapb(cfg *config, args ...string) error {
	if cfg.prevLoactaionAreaURL == nil {
		return errors.New("you're on the first page")
	}
	res, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLoactaionAreaURL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Location areas:")
	for _, area := range res.Results {
		fmt.Printf("- %s\n", area.Name)
	}
	cfg.nextLocationAreaURL = res.Next
	cfg.prevLoactaionAreaURL = res.Previous
	return nil
}