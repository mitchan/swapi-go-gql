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

// func FindFilm(films []types.Film, slice *[]types.Film, filmUrls []string) []types.Film {
// 	if len(filmUrls) == 0 {
// 		return slice
// 	}

// 	for _, url := range filmUrls {
// 		for _, film := range films.Films {
// 			if film.Url == url {
// 				slice = append(slice, film)
// 				break
// 			}
// 		}
// 	}

// 	return slice
// }
