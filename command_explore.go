package main

import (
	"errors"
	"fmt"
)

func commandExplore(area *string, cfg *config) error {
	if area == nil {
		return errors.New("no area passed")
	}
	resp, err := cfg.pokeapiClient.GetLocationDetails(*area)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %v...\n", *area)
	fmt.Println("Found pokemon:")
	for _, pokemon := range resp.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}
	return nil

}
