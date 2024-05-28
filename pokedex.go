package main

import (
	"errors"
	"fmt"

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

func (p *Pokedex) Inspect() error {
	if len(p.pokemon) == 0 {
		return errors.New("You have not caught any pokemon.")
	}
	fmt.Println("Caught pokemon:")
	for _, p := range p.pokemon {
		fmt.Printf(" -%s\n", p.Name)
	}
	return nil
}
