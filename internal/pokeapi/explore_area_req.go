package pokeapi

import (
	"net/http"
	"fmt"
	"io"
	"encoding/json"
)

func (c *Client) ExploreLocationAreas(areaName string) (ExploreLocationArea, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint + "/" + areaName

	dat, ok := c.cache.Get(fullURL)
	if ok {
		exploreLocationArea := ExploreLocationArea{}
		err := json.Unmarshal(dat, &exploreLocationArea)
		if err!=nil {
			return ExploreLocationArea{}, err
		}
		return exploreLocationArea, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)

	if err!=nil {
		return ExploreLocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err!=nil {
		return ExploreLocationArea{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode>399 {
		return ExploreLocationArea{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err!=nil {
		return ExploreLocationArea{}, err
	}

	exploreLocationArea := ExploreLocationArea{}
	err = json.Unmarshal(dat, &exploreLocationArea)
	if err!=nil {
		return ExploreLocationArea{}, err
	}

	c.cache.Add(fullURL, dat)
	return exploreLocationArea, nil
}