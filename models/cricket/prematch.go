package cricket


// CricketPrematchData represents the structure of the cricket prematch JSON
type CricketPrematchData struct {
	Success int `json:"success"`
	Results []struct {
		FI       string `json:"FI"`
		EventID  string `json:"event_id"`
		FirstOver struct {
			UpdatedAt string `json:"updated_at"`
			Key       string `json:"key"`
			SP        struct {
				FirstOverTotalRuns struct {
					ID    string `json:"id"`
					Name  string `json:"name"`
					Odds  []Odd  `json:"odds"`
				} `json:"1st_over_total_runs"`
				FirstOverTotalRunsOddEven struct {
					ID    string `json:"id"`
					Name  string `json:"name"`
					Odds  []Odd  `json:"odds"`
				} `json:"1st_over_total_runs_odd_even"`
			} `json:"sp"`
		} `json:"1st_over"`
		Innings1 struct {
			UpdatedAt string `json:"updated_at"`
			Key       string `json:"key"`
			SP        struct {
				FirstInningsScore struct {
					ID    string `json:"id"`
					Name  string `json:"name"`
					Odds  []Odd  `json:"odds"`
				} `json:"1st_innings_score"`
				FirstInningsOfMatchBowledOut struct {
					ID    string `json:"id"`
					Name  string `json:"name"`
					Odds  []Odd  `json:"odds"`
					Open  int    `json:"open"`
				} `json:"1st_innings_of_match_bowled_out?"`
			} `json:"sp"`
		} `json:"innings_1"`
		Main struct {
			UpdatedAt string `json:"updated_at"`
			Key       string `json:"key"`
			SP        struct {
				ToWinTheMatch struct {
					ID   string `json:"id"`
					Name string `json:"name"`
					Odds []Odd  `json:"odds"`
				} `json:"to_win_the_match"`
				TeamTopBatter struct {
					ID   string `json:"id"`
					Name string `json:"name"`
					Odds []Odd  `json:"odds"`
				} `json:"team_top_batter"`
				TeamTopBowler struct {
					ID   string `json:"id"`
					Name string `json:"name"`
					Odds []Odd  `json:"odds"`
					Open int    `json:"open"`
				} `json:"team_top_bowler"`
				PlayerOfTheMatch struct {
					ID   string `json:"id"`
					Name string `json:"name"`
					Odds []Odd  `json:"odds"`
					Open int    `json:"open"`
				} `json:"player_of_the_match"`
				FirstWicketMethod struct {
					ID   string `json:"id"`
					Name string `json:"name"`
					Odds []Odd  `json:"odds"`
					Open int    `json:"open"`
				} `json:"1st_wicket_method"`
				PlayerPerformance struct {
					ID   string `json:"id"`
					Name string `json:"name"`
					Odds []Odd  `json:"odds"`
				} `json:"player_performance"`
			} `json:"sp"`
		} `json:"main"`
		Match struct {
			UpdatedAt string `json:"updated_at"`
			Key       string `json:"key"`
			SP        struct {
				TeamToMakeHighest1st6OversScore struct {
					ID   string `json:"id"`
					Name string `json:"name"`
					Odds []Odd  `json:"odds"`
				} `json:"team_to_make_highest_1st_6_overs_score"`
				ToGoToSuperOver struct {
					ID   string `json:"id"`
					Name string `json:"name"`
					Odds []Odd  `json:"odds"`
				} `json:"to_go_to_super_over?"`
				RunsAtFallOf1stWicket struct {
					ID   string `json:"id"`
					Name string `json:"name"`
					Odds []Odd  `json:"odds"`
				} `json:"runs_at_fall_of_1st_wicket"`
				AFiftyToBeScored struct {
					ID   string `json:"id"`
					Name string `json:"name"`
					Odds []Odd  `json:"odds"`
				} `json:"a_fifty_to_be_scored"`
				AHundredToBeScored struct {
					ID   string `json:"id"`
					Name string `json:"name"`
					Odds []Odd  `json:"odds"`
				} `json:"a_hundred_to_be_scored_in_the_match"`
				MostMatchSixes struct {
					ID   string `json:"id"`
					Name string `json:"name"`
					Odds []Odd  `json:"odds"`
				} `json:"most_match_sixes"`
				MostMatchFours struct {
					ID   string `json:"id"`
					Name string `json:"name"`
					Odds []Odd  `json:"odds"`
				} `json:"most_match_fours"`
				HighestIndividualScore struct {
					ID   string `json:"id"`
					Name string `json:"name"`
					Odds []Odd  `json:"odds"`
				} `json:"highest_individual_score"`
			} `json:"sp"`
		} `json:"match"`
		Player struct {
			UpdatedAt string `json:"updated_at"`
			Key       string `json:"key"`
			SP        struct {
				BatterMatchRuns struct {
					ID   string `json:"id"`
					Name string `json:"name"`
					Odds []Odd  `json:"odds"`
				} `json:"batter_match_runs"`
				BatterMilestones struct {
					ID   string `json:"id"`
					Name string `json:"name"`
					Odds []Odd  `json:"odds"`
				} `json:"batter_milestones"`
				BowlerTotalMatchWickets struct {
					ID   string `json:"id"`
					Name string `json:"name"`
					Odds []Odd  `json:"odds"`
				} `json:"bowler_total_match_wickets"`
			} `json:"sp"`
		} `json:"player"`
		Schedule struct {
			UpdatedAt string `json:"updated_at"`
			Key       string `json:"key"`
			SP        struct {
				Main []Odd `json:"main"`
			} `json:"sp"`
		} `json:"schedule"`
	} `json:"results"`
}
