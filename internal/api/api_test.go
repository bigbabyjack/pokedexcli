package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Mock response JSON
const mockResponse = `{
	"count": 2,
	"next": "https://pokeapi.co/api/v2/location?offset=2&limit=2",
	"previous": null,
	"results": [
		{
			"name": "kanto",
			"url": "https://pokeapi.co/api/v2/location/1/"
		},
		{
			"name": "johto",
			"url": "https://pokeapi.co/api/v2/location/2/"
		}
	]
}`

// TestMakeGetRequest tests the makeGetRequest function
func TestMakeGetRequest(t *testing.T) {
	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(mockResponse))
	}))
	defer server.Close()

	body, err := makeGetRequest(server.URL)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	var response PokemonLocationRequest
	if err := json.Unmarshal(body, &response); err != nil {
		t.Fatalf("Expected no error in unmarshalling, got %v", err)
	}

	if response.Count != 2 {
		t.Errorf("Expected count 2, got %d", response.Count)
	}
}

// TestGetPokemonLocation tests the GetPokemonLocation function
func TestGetPokemonLocation(t *testing.T) {
	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(mockResponse))
	}))
	defer server.Close()

	response, err := GetPokemonLocation(server.URL)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if response.Count != 2 {
		t.Errorf("Expected count 2, got %d", response.Count)
	}

	if response.Next != "https://pokeapi.co/api/v2/location?offset=2&limit=2" {
		t.Errorf("Expected next URL, got %s", response.Next)
	}

	if len(response.Locations) != 2 {
		t.Fatalf("Expected 2 locations, got %d", len(response.Locations))
	}

	expectedLocations := []struct {
		Name string
		URL  string
	}{
		{"kanto", "https://pokeapi.co/api/v2/location/1/"},
		{"johto", "https://pokeapi.co/api/v2/location/2/"},
	}

	for i, location := range response.Locations {
		if location.Location != expectedLocations[i].Name {
			t.Errorf("Expected location name %s, got %s", expectedLocations[i].Name, location.Location)
		}
		if location.URL != expectedLocations[i].URL {
			t.Errorf("Expected location URL %s, got %s", expectedLocations[i].URL, location.URL)
		}
	}
}
