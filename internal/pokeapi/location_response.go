package pokeapi

type LocationsPageResponse struct {
	Count    int           `json:"count"`
	Next     *string       `json:"next"`
	Previous *string       `json:"previous"`
	Results  []ApiResource `json:"results"`
}

type LocationArea struct {
	ID                   int                   `json:"id"`
	Name                 string                `json:"name"`
	GameIndex            int                   `json:"game_index"`
	EncounterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
	Location             ApiResource           `json:"location"`
	Names                []Name                `json:"names"`
	PokemonEncounters    []PokemonEncounter    `json:"pokemon_encounters"`
}

type EncounterMethodRate struct {
	EncounterMethod ApiResource               `json:"encounter_method"`
	VersionDetails  []EncounterVersionDetails `json:"version_details"`
}

type EncounterVersionDetails struct {
	Rate    int         `json:"rate"`
	Version ApiResource `json:"version"`
}

type Name struct {
	Name     string      `json:"name"`
	Language ApiResource `json:"language"`
}

type PokemonEncounter struct {
	Pokemon        ApiResource              `json:"pokemon"`
	VersionDetails []VersionEncounterDetail `json:"version_details"`
}

type VersionEncounterDetail struct {
	Version          ApiResource `json:"version"`
	MaxChance        int         `json:"max_chance"`
	EncounterDetails []Encounter `json:"encounter_details"`
}

type Encounter struct {
	MinLevel        int                       `json:"min_level"`
	MaxLevel        int                       `json:"max_level"`
	ConditionValues []EncounterConditionValue `json:"condition_values"`
	Chance          int                       `json:"chance"`
	Method          ApiResource               `json:"method"`
}

type EncounterConditionValue struct {
	ID        int         `json:"id"`
	Name      string      `json:"name"`
	Condition ApiResource `json:"condition"`
	Names     []Name      `json:"names"`
}

type ApiResource struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
