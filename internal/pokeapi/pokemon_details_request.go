package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemonDetails(name *string) (PokemonDetails, error) {
	endpoint := "/pokemon/"
	fullURL := baseURL + endpoint + *name

	dat, ok := c.cache.Get(fullURL)
	if ok {
		pokemonDetails := PokemonDetails{}
		err := json.Unmarshal(dat, &pokemonDetails)
		if err != nil {
			return PokemonDetails{}, err
		}
		return pokemonDetails, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return PokemonDetails{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonDetails{}, err
	}
	pokemonDetails := PokemonDetails{}
	dat, err = io.ReadAll(resp.Body)
	c.cache.Add(fullURL, dat)
	err = json.Unmarshal(dat, &pokemonDetails)
	if err != nil {
		return PokemonDetails{}, err
	}

	return pokemonDetails, nil
}
