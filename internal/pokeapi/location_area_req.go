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

	// Check if the data is in the cache
	if cachedData, found := c.cache.Get(fullURL); found {
		fmt.Println("Cache hit!")
		var locationAreaRes LocationAreaRes
		err := json.Unmarshal(cachedData, &locationAreaRes)
		if err != nil {
			return LocationAreaRes{}, err
		}
		return locationAreaRes, nil
	}

	fmt.Println("Cache miss!")

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

	// Add the data to the cache
	c.cache.Add(fullURL, data)

	locationAreaRes := LocationAreaRes{}
	err = json.Unmarshal(data, &locationAreaRes)
	if err != nil {
		return LocationAreaRes{}, err
	}

	return locationAreaRes, nil
}