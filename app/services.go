package app

import (
	"io"
	"net/http"
)

const url = "https://api.coingecko.com/api/v3/"

// Make a http request to the given url and return the response
func GetCoinsData() ([]byte, error) {
	res, err := http.Get(url + "coins/markets?vs_currency=usd&order=market_cap_desc&per_page=100&page=1")

	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}
