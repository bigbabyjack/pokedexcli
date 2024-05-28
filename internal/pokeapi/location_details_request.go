package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationDetails(area string) (LocationDetailsResponse, error) {
	endpoint := "/location-area/"
	fullURL := baseURL + endpoint + area

	// cache check
	if dat, ok := c.cache.Get(fullURL); ok {
		locationDetailsResponse := LocationDetailsResponse{}
		err := json.Unmarshal(dat, &locationDetailsResponse)
		if err != nil {
			return LocationDetailsResponse{}, err
		}

		return locationDetailsResponse, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationDetailsResponse{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationDetailsResponse{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationDetailsResponse{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	dat, err := io.ReadAll(resp.Body)
	c.cache.Add(fullURL, dat)
	if err != nil {
		return LocationDetailsResponse{}, err
	}

	locationDetailsResponse := LocationDetailsResponse{}
	err = json.Unmarshal(dat, &locationDetailsResponse)
	if err != nil {
		return LocationDetailsResponse{}, err
	}

	return locationDetailsResponse, nil

}
