package pokeapi

import (
	"fmt"
	"io"
)

func (c *Client) FetchLocations(url string) ([]byte, error) {
	res, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode > 299 {
		return nil, fmt.Errorf("bad status code: %d", res.StatusCode)
	}

	return body, err
}
