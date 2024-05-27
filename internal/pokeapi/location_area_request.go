package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResponse, error) {
	endpoint := "/location-area?offset=0&limit=20"
	fullURL := baseURL + endpoint

	if pageURL != nil {
		fullURL = *pageURL
	}

	if dat, ok := c.cache.Get(fullURL); ok {
		locationAreasResponse := LocationAreasResponse{}
		err := json.Unmarshal(dat, &locationAreasResponse)
		if err != nil {
			return LocationAreasResponse{}, err
		}

		return locationAreasResponse, nil
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
	c.cache.Add(fullURL, dat)
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
