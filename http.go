package main

import (
	"fmt"
	"io"
	"net/http"
)

func FetchLocations(url string) ([]byte, error) {
	res, err := http.Get(url)
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
