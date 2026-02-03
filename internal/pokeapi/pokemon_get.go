package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName
	if val, ok := c.cache.Get(url); ok {
		var pokemonResp Pokemon
		if err := json.Unmarshal(val, &pokemonResp); err != nil {
			return Pokemon{}, err
		}
		return pokemonResp, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return Pokemon{}, fmt.Errorf("bad status code: %d", resp.StatusCode)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	var pokemonResp Pokemon
	if err := json.Unmarshal(dat, &pokemonResp); err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, dat)

	return pokemonResp, nil
}
