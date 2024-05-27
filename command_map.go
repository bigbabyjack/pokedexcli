package main

import (
	"errors"
	"fmt"
	"log"
)

func commandMap(a *string, cfg *config) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.NextLocationURL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Location Areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	cfg.NextLocationURL = resp.Next
	cfg.PreviousLocationURL = resp.Previous
	return nil

}

func commandMapb(a *string, cfg *config) error {
	if cfg.PreviousLocationURL == nil {
		return errors.New("You are already at the first locations.")
	}

	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.PreviousLocationURL)
	if err != nil {
		return fmt.Errorf("Error getting locations: %v", err)
	}

	fmt.Println("Location Areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	cfg.NextLocationURL = resp.Next
	cfg.PreviousLocationURL = resp.Previous
	return nil

}
