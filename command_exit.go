package main

import (
	"fmt"
	"os"
)

func commandExit(a *string, config *config) error {
	fmt.Println("Exiting Pokedex. Goodbye!")
	os.Exit(0)
	return nil
}
