package requests

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/Kenedy228/pokedex/internal/cache"
	"github.com/Kenedy228/pokedex/internal/entities"
)

const cacheLifetime = 30 * time.Second
const URL = "https://pokeapi.co/api/v2/location/?offset=80&limit=20"

type Mapper struct {
	URL   string
	cache cache.Cache
}

type Response struct {
	Links []Link `json:"results"`
}

type Link struct {
	URL string `json:"url"`
}

func NewMapper() (*Mapper, error) {
	mapper := &Mapper{
		URL:   URL,
		cache: cache.NewCache(cacheLifetime),
	}

	return mapper, nil
}

func (m *Mapper) GetAreas() (map[string]entities.Area, error) {
	var data []byte

	cached, ok := m.cache.Get(m.URL)

	if ok {
		data = cached
	} else {
		res, err := http.Get(m.URL)

		if err != nil {
			return nil, err
		}

		data, err = io.ReadAll(res.Body)
		defer res.Body.Close()

		if err != nil {
			return nil, err
		}
	}

	var response Response

	if err := json.Unmarshal(data, &response); err != nil {
		return nil, err
	}

	locations := m.handleLocations(response.Links)

	areas := map[string]entities.Area{}

	for _, v := range locations {
		for _, area := range v.Areas {
			areas[area.Name] = area
		}
	}

	return areas, nil
}

func (m *Mapper) handleLocations(links []Link) []entities.Location {
	result := []entities.Location{}

	var wg sync.WaitGroup
	var mux sync.Mutex

	for _, v := range links {
		wg.Add(1)

		go func(url string) {
			defer wg.Done()

			var data []byte

			mux.Lock()
			cached, isPresent := m.cache.Get(url)
			mux.Unlock()

			if isPresent {
				data = cached
			} else {
				res, err := http.Get(url)
				if err != nil {
					return
				}

				data, err = io.ReadAll(res.Body)
				defer res.Body.Close()

				if err != nil {
					return
				}

				mux.Lock()
				m.cache.Add(url, data)
				mux.Unlock()
			}

			var location entities.Location

			if err := json.Unmarshal(data, &location); err != nil {
				return
			}

			mux.Lock()
			result = append(result, location)
			mux.Unlock()

		}(v.URL)
	}

	wg.Wait()

	return result
}

func (m *Mapper) GetPocemonsByArea(name string) ([]entities.Pokemon, error) {
	areas, err := m.GetAreas()

	if err != nil {
		return nil, err
	}

	a, ok := areas[name]

	if !ok {
		return nil, fmt.Errorf("No area found")
	}

	encounter, err := m.findPokemonEncounter(a.URL)

	if err != nil {
		return nil, err
	}

	pokemons := []entities.Pokemon{}

	for _, r := range encounter.Pokemons {
		pokemons = append(pokemons, r.Pokemon)
	}

	return pokemons, nil
}

func (m *Mapper) findPokemonEncounter(url string) (*entities.PokemonEncounter, error) {
	var data []byte

	cached, ok := m.cache.Get(url)

	if ok {
		data = cached
	} else {
		res, err := http.Get(url)

		if err != nil {
			return nil, err
		}

		data, err = io.ReadAll(res.Body)
		defer res.Body.Close()

		if err != nil {
			return nil, err
		}
	}

	var encounter entities.PokemonEncounter

	if err := json.Unmarshal(data, &encounter); err != nil {
		return nil, err
	}

	return &encounter, nil
}

func (m *Mapper) FindPokemonExperience(name string) (*entities.PokemonStats, error) {
	areas, err := m.GetAreas()

	if err != nil {
		return nil, err
	}

	var url string

	for _, a := range areas {
		encounter, err := m.findPokemonEncounter(a.URL)

		if err != nil {
			return nil, err
		}

		for _, p := range encounter.Pokemons {
			if p.Pokemon.Name == name {
				url = p.Pokemon.URL
				break
			}
		}
	}

	if url == "" {
		return nil, fmt.Errorf("Not Found pokemon with name %s", name)
	}

	var data []byte

	cached, ok := m.cache.Get(url)

	if ok {
		data = cached
	} else {
		res, err := http.Get(url)

		if err != nil {
			return nil, err
		}

		data, err = io.ReadAll(res.Body)
		defer res.Body.Close()

		if err != nil {
			return nil, err
		}
	}

	var stats entities.PokemonStats

	err = json.Unmarshal(data, &stats)

	if err != nil {
		return nil, err
	}

	return &stats, nil
}
