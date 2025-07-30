package entities

type PokemonEncounter struct {
	Pokemons []PokemonResponse `json:"pokemon_encounters"`
}

type PokemonResponse struct {
	Pokemon Pokemon `json:"pokemon"`
}

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokemonStats struct {
	Experience int `json:"base_experience"`
}
