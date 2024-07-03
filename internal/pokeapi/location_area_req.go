package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreaRes, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint
	if pageURL != nil {
		fullURL = *pageURL
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreaRes{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaRes{}, err
	}

	defer res.Body.Close()

	if res.StatusCode > 399 {
		return LocationAreaRes{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaRes{}, err
	}

	locationAreaRes := LocationAreaRes{}
	err = json.Unmarshal(data, &locationAreaRes)
	if err != nil {
		return LocationAreaRes{}, err
	}

	return locationAreaRes, nil
}