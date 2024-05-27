package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResponse, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint

	if pageURL != nil {
		fullURL = *pageURL
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreasResponse{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	locationAreasResponse := LocationAreasResponse{}
	err = json.Unmarshal(dat, &locationAreasResponse)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	return locationAreasResponse, nil
}

// func (c *Client) makeGetRequest(s string) ([]byte, error) {
// 	resp, err := http.Get(s)
// 	if err != nil {
// 		return nil, errors.New("Error in GET request.")
// 	}
// 	defer resp.Body.Close()
// 	body, err := io.ReadAll(resp.Body)
// 	return body, nil
//
// }
//
// func (c *Client) GetPokemonLocation(url string) (LocationAreasRequest, error) {
// 	body, err := c.makeGetRequest(url)
// 	var request LocationAreasRequest
// 	if err != nil {
// 		return request, fmt.Errorf("[FAILURE] failed GET request: %w", err)
// 	}
//
// 	if err := json.Unmarshal(body, &request); err != nil {
// 		return request, fmt.Errorf("failed to parse JSON: %w", err)
// 	}
//
// 	return request, nil
// }