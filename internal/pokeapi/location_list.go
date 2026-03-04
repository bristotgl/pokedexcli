package pokeapi

import (
	"encoding/json"
	"net/http"

	"github.com/bristotgl/pokedexcli/internal/converter"
)

func (c *Client) ListLocations(pageURL *string) (LocationsPage, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if cachedPageBytes, ok := c.cache.Get(url); ok {
		return converter.FromBytes[LocationsPage](cachedPageBytes)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationsPage{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationsPage{}, err
	}
	defer resp.Body.Close()

	var locationsPage LocationsPage
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&locationsPage); err != nil {
		return LocationsPage{}, err
	}

	pageBytes, err := converter.ToBytes(locationsPage)
	if err != nil {
		return LocationsPage{}, err
	}
	c.cache.Add(url, pageBytes)

	return locationsPage, nil
}
