package pokeapi

import (
	"encoding/json"
	"net/http"

	"github.com/bristotgl/pokedexcli/internal/converter"
)

func (c *Client) ListLocations(pageURL *string) (LocationsPageResponse, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if cachedPageBytes, ok := c.cache.Get(url); ok {
		return converter.FromBytes[LocationsPageResponse](cachedPageBytes)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationsPageResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationsPageResponse{}, err
	}
	defer resp.Body.Close()

	var locationsPage LocationsPageResponse
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&locationsPage); err != nil {
		return LocationsPageResponse{}, err
	}

	pageBytes, err := converter.ToBytes(locationsPage)
	if err != nil {
		return LocationsPageResponse{}, err
	}
	c.cache.Add(url, pageBytes)

	return locationsPage, nil
}
