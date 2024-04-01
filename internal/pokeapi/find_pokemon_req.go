package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"fmt"
)

func (c *Client) FindPokemon(pokemonName string) (Pokemon, error) {
	endpoint := "/pokemon/" + pokemonName
	fullURL := baseURL + endpoint

	dat, ok := c.cache.Get(fullURL)
	if ok {
		pokemonObj := Pokemon{}
		err := json.Unmarshal(dat, &pokemonObj)
		if err!=nil {
			return Pokemon{}, err
		}
		return pokemonObj, nil
	}


	req, err := http.NewRequest("GET", fullURL, nil)

	if err!=nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err!=nil {
		return Pokemon{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode>399 {
		return Pokemon{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)

	if err!=nil {
		return Pokemon{}, err
	}

	pokemonObj := Pokemon{}
	err = json.Unmarshal(dat, &pokemonObj)

	if err!=nil {
		return Pokemon{}, err
	}

	c.cache.Add(fullURL, dat)
	return pokemonObj, nil
} 