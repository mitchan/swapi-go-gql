package utils

import (
	"io/ioutil"
	"net/http"
)

const baseUrl = "https://swapi.dev/api/"

func LoadData(endpoint string) ([]byte, error) {
	resp, err := http.Get(baseUrl + endpoint)
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
