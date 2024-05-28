package main

import (
	"github.com/bigbabyjack/pokedexcli/internal/pokeapi"
)

type Pokemon struct {
	name string
}

type Pokedex struct {
	pokemon map[string]pokeapi.PokemonDetails
}

func NewPokedex() *Pokedex {
	return &Pokedex{
		pokemon: make(map[string]pokeapi.PokemonDetails),
	}
}

func (p *Pokedex) Get(key string) (pokeapi.PokemonDetails, bool) {
	dat, ok := p.pokemon[key]
	if !ok {
		return pokeapi.PokemonDetails{}, false
	}

	return dat, true
}

func (p *Pokedex) Add(key string, dat pokeapi.PokemonDetails) error {
	p.pokemon[key] = dat
	return nil
}

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
