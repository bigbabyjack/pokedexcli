package main

import (
	"fmt"
	"os"
)

func (r *Repl) commandExit(config *Config) error {
	fmt.Println("Exiting Pokedex. Goodbye!")
	os.Exit(0)
	return nil
}
