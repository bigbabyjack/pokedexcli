package main

import (
	"fmt"
	"math/rand"
)

const maxCatchTreshold float64 = 0.3
const baseCatchTreshold float64 = 0.3

func calculateCatchThreshold(xp int) float64 {
	calculatedCatchThreshold := baseCatchTreshold + ((0.5 * float64(xp)) / (1 + float64(xp)))
	return min(maxCatchTreshold, calculatedCatchThreshold)
}

func commandCatch(name *string, cfg *config) error {
	details, err := cfg.pokeapiClient.GetPokemonDetails(name)
	if err != nil {
		return err
	}

	_, ok := cfg.pokedex.Get(*name)
	if ok {
		return fmt.Errorf("Pokemon: %s already owned.", *name)
	}

	thresh := calculateCatchThreshold(details.BaseExperience)
	roll := rand.Float64()
	fmt.Printf("Throwing a pokeball at %s...\n", *name)
	if roll > thresh {
		cfg.pokedex.Add(*name, details)
		fmt.Printf("%s was caught!\n", *name)
		return nil
	}
	fmt.Printf("%s escaped!\n", *name)
	return nil

}
