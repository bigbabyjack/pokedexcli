package main

import (
	"fmt"
)

func commandHelp(a *string, config *config) error {
	fmt.Println("Welcome to Pokedex!")
	fmt.Println("Usage:")

	for _, cmd := range getCommands() {
		fmt.Println(cmd.name, ": ", cmd.description)
	}

	return nil
}
