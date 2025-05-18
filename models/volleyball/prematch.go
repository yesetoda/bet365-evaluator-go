package volleyball


// Prematch Data Structures
type PrematchData struct {
	Success int              `json:"success"`
	Results []PrematchResult `json:"results"`
}

type PrematchResult struct {
	FI       string       `json:"FI"`
	EventID  string       `json:"event_id"`
	Main     MainData     `json:"main"`
	Others   []OtherData  `json:"others"`
	Schedule ScheduleData `json:"schedule"`
}

type MainData struct {
	UpdatedAt string `json:"updated_at"`
	Key       string `json:"key"`
	Sp        SpData `json:"sp"`
}

type OtherData struct {
	UpdatedAt string `json:"updated_at"`
	Sp        SpData `json:"sp"`
}

type SpData struct {
	GameLines            MarketData `json:"game_lines"`
	CorrectSetScore      MarketData `json:"correct_set_score"`
	MatchTotalOddEven    MarketData `json:"match_total_odd_even"`
	Set1Lines            MarketData `json:"set_1_lines"`
	Set1ToGoToExtraPoints MarketData `json:"set_1_to_go_to_extra_points"`
	Set1TotalOddEven     MarketData `json:"set_1_total_odd_even"`
}

type MarketData struct {
	ID    string     `json:"id"`
	Name  string     `json:"name"`
	Odds  []OddsData `json:"odds"`
	Open  int        `json:"open,omitempty"`
}

type OddsData struct {
	ID       string `json:"id"`
	Odds     string `json:"odds"`
	Name     string `json:"name"`
	Header   string `json:"header"`
	Handicap string `json:"handicap,omitempty"`
}

type ScheduleData struct {
	UpdatedAt string     `json:"updated_at"`
	Key       string     `json:"key"`
	Sp        ScheduleSp `json:"sp"`
}

type ScheduleSp struct {
	Main []ScheduleOddsData `json:"main"`
}

type ScheduleOddsData struct {
	ID       string `json:"id"`
	Odds     string `json:"odds"`
	Name     string `json:"name"`
	Handicap string `json:"handicap,omitempty"`
}
