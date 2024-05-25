package main

import (
	"errors"
	"fmt"
)

func (r *Repl) commandHelp() error {
	fmt.Println("Welcome to Pokedex!")
	fmt.Println("Usage:")

	if len(r.commands) == 0 {
		return errors.New("Repl has no commands!")
	}

	for c := range r.commands {
		fmt.Println(c, ": ", r.commands[c].description)
	}

	return nil
}
