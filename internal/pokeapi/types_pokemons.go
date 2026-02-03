package pokeapi

type Pokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		Stat struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"stat"`
		Effort   int `json:"effort"`
		BaseStat int `json:"base_stat"`
	}
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"type"`
	}
}
