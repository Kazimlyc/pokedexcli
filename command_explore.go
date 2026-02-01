package main

// todo
func commandExplore(cfg *config, args ...string) error {
	type LocationAreasResp struct {
		Count    int
		Next     *string
		Previous *string
		Results  []struct {
			Name string
			URL  string
		}
	}

	return nil
}
