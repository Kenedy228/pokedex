package requests

import (
	"encoding/json"
	"github.com/Kenedy228/pokedex/entities"
	"net/http"
)

type Mapper struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		URL string `json:"url"`
	} `json:"results"`
}

func NewMapper(URL string) (*Mapper, error) {
	mapper := &Mapper{}

	res, err := http.Get(URL)

	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(res.Body)
	defer res.Body.Close()

	err = decoder.Decode(mapper)

	if err != nil {
		return nil, err
	}

	return mapper, nil
}

func (m *Mapper) Handle() ([]entities.Location, error) {
	results := []entities.Location{}

	for _, v := range m.Results {
		res, err := http.Get(v.URL)

		if err != nil {
			return nil, err
		}

		decoder := json.NewDecoder(res.Body)
		defer res.Body.Close()

		var location entities.Location

		err = decoder.Decode(&location)

		if err != nil {
			return nil, err
		}

		results = append(results, location)
	}

	return results, nil
}
