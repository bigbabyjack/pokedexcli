package main

import (
	"fmt"
	"os"
)

func commandExit(config *config) error {
	fmt.Println("Exiting Pokedex. Goodbye!")
	os.Exit(0)
	return nil
}
