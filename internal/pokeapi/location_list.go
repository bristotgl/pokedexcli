package pokeapi

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"net/http"

	"github.com/bristotgl/pokedexcli/internal/pokecache"
)

func (c *Client) ListLocations(pageURL *string, pokeCache *pokecache.Cache) (LocationsPage, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	cachedPageBytes, exists := pokeCache.Get(url)
	if exists {
		return convertBytesToPage(cachedPageBytes)
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

	pageBytes, err := convertPageToBytes(locationsPage)
	if err != nil {
		return LocationsPage{}, err
	}
	pokeCache.Add(url, pageBytes)

	return locationsPage, nil
}

func convertBytesToPage(cachedPageBytes []byte) (LocationsPage, error) {
	var cachedPage LocationsPage

	decoder := gob.NewDecoder(bytes.NewReader(cachedPageBytes))
	if err := decoder.Decode(&cachedPage); err != nil {
		return LocationsPage{}, err
	}

	return cachedPage, nil
}

func convertPageToBytes(locationsPage LocationsPage) ([]byte, error) {
	var buf bytes.Buffer

	encoder := gob.NewEncoder(&buf)
	if err := encoder.Encode(locationsPage); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
