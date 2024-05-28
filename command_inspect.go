package main

import (
	"errors"
)

func commandInspect(name *string, cfg *config) error {
	dat, ok := cfg.pokedex.Get(*name)
	if !ok {
		return errors.New("You have not caught that pokemon.")
	}
	err := dat.Inspect()
	if err != nil {
		return err
	}
	return nil
}
