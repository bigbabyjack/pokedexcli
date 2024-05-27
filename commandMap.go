package main

import (
	"errors"
	"fmt"
	"github.com/bigbabyjack/pokedexcli/internal/api"
)

func (r *Repl) commandMap(config *Config) error {
	// base URL for location API
	url := "https://pokeapi.co/api/v2/location"

	//if we already have a next page to get, go to it
	if config.Next != "" {
		url = config.Next
	}
	locations, ok := r.cache.Get(url)
	if ok {
		fmt.Println("Found cached map.")
		for _, l := range locations {
			fmt.Println(l)
		}
		return nil
	}
	fmt.Println("Fetching from API")

	// if it isn't in the cache, get it from the api
	fmt.Println("Requesting api.")
	request, err := api.GetPokemonLocation(url)
	if err != nil {
		return err
	}

	for _, location := range request.Locations {
		fmt.Println(location.Location)
		r.cache.Add(url, []byte(location.Location))
	}
	// update the config
	config.Next = request.Next
	config.Previous = request.Previous

	// print out the locations
	return nil
}

func (r *Repl) commandMapb(config *Config) error {
	if config.Previous == "" {
		return errors.New("You are on the first page of locations.")
	}
	url := config.Previous

	// check if the url is in the cache and return if it is
	locations, ok := r.cache.Get(url)
	if ok {
		fmt.Println("Found cached map.")
		for _, l := range locations {
			fmt.Println(l)
		}
		return nil
	}
	fmt.Println("Fetching from API")

	request, err := api.GetPokemonLocation(url)
	if err != nil {
		return errors.New("Unable to get locations.")
	}
	config.Next = request.Next
	config.Previous = request.Previous
	for _, location := range request.Locations {
		fmt.Println(location.Location)
		r.cache.Add(url, []byte(location.Location))
	}
	return nil
}
