package main

import (
	"github.com/bigbabyjack/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient       pokeapi.Client
	NextLocationURL     *string
	PreviousLocationURL *string
	pokedex             *Pokedex
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(),
		pokedex:       NewPokedex(),
	}
	runRepl(&cfg)

}
