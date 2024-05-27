package main

import "github.com/bigbabyjack/pokedexcli/internal/pokeapi"

type config struct {
	pokeapiClient       pokeapi.Client
	NextLocationURL     *string
	PreviousLocationURL *string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(),
	}
	runRepl(&cfg)

}
