package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type PokemonLocationRequest struct {
	Count     int    `json:"count"`
	Next      string `json:"next"`
	Previous  string `json:"previous"`
	Locations []struct {
		Location string `json:"name"`
		URL      string `json:"url"`
	} `json:"results"`
}

func makeGetRequest(s string) ([]byte, error) {
	resp, err := http.Get(s)
	if err != nil {
		return nil, errors.New("Error in GET request.")
	}
	defer resp.Body.Close()
	fmt.Println("Response status:", resp.Status)
	body, err := io.ReadAll(resp.Body)
	return body, nil

}

func GetPokemonLocation(url string) (PokemonLocationRequest, error) {
	body, err := makeGetRequest(url)
	var request PokemonLocationRequest
	if err != nil {
		return request, fmt.Errorf("[FAILURE] failed GET request: %w", err)
	}

	if err := json.Unmarshal(body, &request); err != nil {
		return request, fmt.Errorf("failed to parse JSON: %w", err)
	}

	fmt.Printf("Count: %d\n", request.Count)
	fmt.Printf("Next: %s\n", request.Next)
	fmt.Printf("Previous: %s\n", request.Previous)
	for _, location := range request.Locations {
		fmt.Printf("Location: %s, URL: %s\n", location.Location, location.URL)
	}

	return request, nil
}
