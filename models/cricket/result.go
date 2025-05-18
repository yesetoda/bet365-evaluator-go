package cricket

// CricketResultData represents the structure of the cricket result JSON
type CricketResultData struct {
	Success int `json:"success"`
	Results []struct {
		ID         string `json:"id"`
		SportID    string `json:"sport_id"`
		TimeStatus string `json:"time_status"`
		League     struct {
			ID   string `json:"id"`
			Name string `json:"name"`
			CC   string `json:"cc"`
		} `json:"league"`
		Home struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			ImageID string `json:"image_id"`
			CC      string `json:"cc"`
		} `json:"home"`
		Away struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			ImageID string `json:"image_id"`
			CC      string `json:"cc"`
		} `json:"away"`
		SS    string `json:"ss"`
		Extra struct {
			StadiumData struct {
				ID        string `json:"id"`
				Name      string `json:"name"`
				City      string `json:"city"`
				Country   string `json:"country"`
				Capacity  string `json:"capacity"`
				GoogleCoo string `json:"googlecoords"`
			} `json:"stadium_data"`
		} `json:"extra"`
		HasLineup         int    `json:"has_lineup"`
		InplayCreatedAt   string `json:"inplay_created_at"`
		InplayUpdatedAt   string `json:"inplay_updated_at"`
		ConfirmedAt       string `json:"confirmed_at"`
		Bet365ID          string `json:"bet365_id"`
	} `json:"results"`
}