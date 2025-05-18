package cricket


// Odd represents a betting odd from the JSON
type Odd struct {
	ID       string `json:"id"`
	Odds     string `json:"odds"`
	Name     string `json:"name"`
	Name2    string `json:"name2,omitempty"`
	Header   string `json:"header,omitempty"`
	Handicap string `json:"handicap,omitempty"`
}

// BetSelection represents a selected bet
type BetSelection struct {
	Market       string
	Selection    string
	Odds         float64
	IsWinner     bool
	Evaluation   string
	ConfidenceLevel string
	AvailableOptions []string
	MarketDescription string
	PotentialProfit float64
	RiskAssessment string
	OddsDecimal    float64
	OddsAmerican   string
	OddsFractional string
}

// DetailedMatchInfo contains enriched match information
type DetailedMatchInfo struct {
	HomeTeam      string
	AwayTeam      string
	HomeScore     int
	AwayScore     int
	Stadium       string
	City          string
	Country       string
	Capacity      string
	MatchDate     string
	LeagueName    string
	BattingStats  map[string]BattingStats
	BowlingStats  map[string]BowlingStats
}

// BattingStats represents key batting statistics for demonstration
type BattingStats struct {
	Runs         int
	Balls        int
	StrikeRate   float64
	Boundaries   int
	Sixes        int
}

// BowlingStats represents key bowling statistics for demonstration
type BowlingStats struct {
	Overs        float64
	RunsConceded int
	Wickets      int
	Economy      float64
}

// BettingHistory represents simulated past betting performance
type BettingHistory struct {
	Market       string
	WinPercentage float64
	AvgOdds      float64
	TotalBets    int
	ProfitLoss   float64
}