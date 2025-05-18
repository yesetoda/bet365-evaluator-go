package volleyball
// Result Data Structures
type ResultData struct {
	Success int           `json:"success"`
	Results []MatchResult `json:"results"`
}

type MatchResult struct {
	ID              string        `json:"id"`
	SportID         string        `json:"sport_id"`
	Time            string        `json:"time"`
	TimeStatus      string        `json:"time_status"`
	League          LeagueInfo    `json:"league"`
	Home            TeamInfo      `json:"home"`
	Away            TeamInfo      `json:"away"`
	SS              string        `json:"ss"`
	Scores          ScoresInfo    `json:"scores"`
	Stats           StatsInfo     `json:"stats"`
	Events          []EventInfo   `json:"events"`
	Extra           ExtraInfo     `json:"extra"`
	InplayCreatedAt string        `json:"inplay_created_at"`
	InplayUpdatedAt string        `json:"inplay_updated_at"`
	ConfirmedAt     string        `json:"confirmed_at"`
	Bet365ID        string        `json:"bet365_id"`
}

type LeagueInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	CC   string `json:"cc"`
}

type TeamInfo struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	ImageID string `json:"image_id"`
	CC      string `json:"cc"`
}

type ScoresInfo struct {
	Set1 SetScore `json:"1"`
	Set2 SetScore `json:"2"`
	Set3 SetScore `json:"3"`
	Set4 SetScore `json:"4"`
	Set5 SetScore `json:"5"`
}

type SetScore struct {
	Home string `json:"home"`
	Away string `json:"away"`
}

type StatsInfo struct {
	PointsWonOnServe []string `json:"points_won_on_serve"`
	LongestStreak    []string `json:"longest_streak"`
}

type EventInfo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

type ExtraInfo struct {
	HomePos    string `json:"home_pos"`
	AwayPos    string `json:"away_pos"`
	BestOfSets string `json:"bestofsets"`
	Round      string `json:"round"`
}

// BetSelection represents a selection made in pre-match
type BetSelection struct {
	Market      string
	MarketID    string
	Selection   string
	SelectionID string
	Odds        float64
	Handicap    string
	StakeAmount float64 // Added stake amount for bet simulation
}

// EvaluationResult represents the result of a bet evaluation
type EvaluationResult struct {
	BetSelection     BetSelection
	IsWin            bool
	Explanation      string
	ProfitLoss       float64 // Added profit/loss calculation
	ReturnAmount     float64 // Added return amount calculation
	ImpliedProbability float64 // Added implied probability
}

// MatchStatistics represents key statistics from the match
type MatchStatistics struct {
	TotalMatchPoints   int
	TotalSet1Points    int
	TotalSet2Points    int
	TotalSet3Points    int
	TotalSet4Points    int
	TotalSet5Points    int
	HomeSetWins        int
	AwaySetWins        int
	MatchWinner        string
	Set1Winner         string
	Set2Winner         string
	Set3Winner         string
	Set4Winner         string
	Set5Winner         string
	Set1ExtraPoints    bool
	Set2ExtraPoints    bool
	Set3ExtraPoints    bool
	Set4ExtraPoints    bool
	Set5ExtraPoints    bool
	CorrectSetScore    string
	TotalSets          int
	MaximumSets        int
}
