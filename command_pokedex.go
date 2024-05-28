package main

func commandPokedex(arg *string, cfg *config) error {
	err := cfg.pokedex.Inspect()
	if err != nil {
		return err
	}

	return nil
}
