package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	endpoint := fmt.Sprintf("/pokemon/%s", name)
	fullURL := baseURL + endpoint

	// Check if the data is in the cache
	if cachedData, found := c.cache.Get(fullURL); found {
		fmt.Println("Cache hit!")
		var pokemonCatch Pokemon
		err := json.Unmarshal(cachedData, &pokemonCatch)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemonCatch, nil
	}

	fmt.Println("Cache miss!")

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	// Add the data to the cache
	c.cache.Add(fullURL, data)

	var pokemonCatch Pokemon
	err = json.Unmarshal(data, &pokemonCatch)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemonCatch, nil
}