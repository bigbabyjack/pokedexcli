package main

import (
	"errors"
	"fmt"
	"github.com/bigbabyjack/pokedexcli/internal/api"
)

func (r *Repl) commandMap(config *Config) error {
	url := "https://pokeapi.co/api/v2/location"
	if config.Next != "" {
		url = config.Next
	}
	request, err := api.GetPokemonLocation(url)
	if err != nil {
		return err
	}
	config.Next = request.Next
	config.Previous = request.Previous
	for _, location := range request.Locations {
		fmt.Println(location.Location)
	}
	return nil
}

func (r *Repl) commandMapb(config *Config) error {
	if config.Previous == "" {
		return errors.New("You are on the first page of locations.")
	}

	request, err := api.GetPokemonLocation(config.Previous)
	if err != nil {
		return errors.New("Unable to get locations.")
	}
	config.Next = request.Next
	config.Previous = request.Previous
	for _, location := range request.Locations {
		fmt.Println(location.Location)
	}
	return nil
}
