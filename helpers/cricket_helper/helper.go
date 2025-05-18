package cricket_helper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
	"github.com/yesetoda/bet365-evaluator-go/models/cricket"
)

// LoadResultData loads cricket result data from a JSON file
func LoadCricketResultData(filename string) (cricket.CricketResultData, error) {
	var data cricket.CricketResultData

	log.Printf("Loading match result data from %s", filename)
	fileContent, err := ioutil.ReadFile(filename)
	if err != nil {
		return data, fmt.Errorf("error reading result file: %v", err)
	}

	if err := json.Unmarshal(fileContent, &data); err != nil {
		return data, fmt.Errorf("error unmarshaling result data: %v", err)
	}

	log.Printf("Successfully loaded result data, found %d results", len(data.Results))
	return data, nil
}

// loadPrematchData loads cricket prematch data from a JSON file
func LoadCricketPrematchData(filename string) (cricket.CricketPrematchData, error) {
	var data cricket.CricketPrematchData

	log.Printf("Loading prematch betting data from %s", filename)
	fileContent, err := ioutil.ReadFile(filename)
	if err != nil {
		return data, fmt.Errorf("error reading prematch file: %v", err)
	}

	if err := json.Unmarshal(fileContent, &data); err != nil {
		return data, fmt.Errorf("error unmarshaling prematch data: %v", err)
	}

	log.Printf("Successfully loaded prematch data, found %d results", len(data.Results))
	return data, nil
}

// extractDetailedMatchInfo extracts comprehensive match information
func ExtractDetailedMatchInfo(data cricket.CricketResultData) cricket.DetailedMatchInfo {
	var info cricket.DetailedMatchInfo

	if len(data.Results) == 0 {
		log.Println("No match results found in data")
		return info
	}

	result := data.Results[0]
	info.HomeTeam = result.Home.Name
	info.AwayTeam = result.Away.Name
	info.Stadium = result.Extra.StadiumData.Name
	info.City = result.Extra.StadiumData.City
	info.Country = result.Extra.StadiumData.Country
	info.Capacity = result.Extra.StadiumData.Capacity
	info.LeagueName = result.League.Name

	// Parse confirmed time if available
	if result.ConfirmedAt != "" {
		// Assuming ISO 8601 format
		t, err := time.Parse(time.RFC3339, result.ConfirmedAt)
		if err == nil {
			info.MatchDate = t.Format("Monday, January 2, 2006")
		} else {
			info.MatchDate = "Date information unavailable"
			log.Printf("Error parsing match date: %v", err)
		}
	}

	// Parse the score
	homeScore, awayScore, err := ParseScore(result.SS)
	if err != nil {
		log.Printf("Failed to parse score: %v", err)
	} else {
		info.HomeScore = homeScore
		info.AwayScore = awayScore
	}

	// Initialize empty stats maps
	info.BattingStats = make(map[string]cricket.BattingStats)
	info.BowlingStats = make(map[string]cricket.BowlingStats)

	return info
}

// parseScore parses the score string into home and away scores
func ParseScore(scoreStr string) (int, int, error) {
	// In cricket, the score format is typically "X-Y" where X is home team score and Y is away team score
	parts := strings.Split(scoreStr, "-")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid score format: %s", scoreStr)
	}

	homeScore, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid home score: %s", parts[0])
	}

	awayScore, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid away score: %s", parts[1])
	}

	return homeScore, awayScore, nil
}

// simulateDetailedMatchStats generates simulated player stats for demonstration
func SimulateDetailedMatchStats(info *cricket.DetailedMatchInfo) {
	// Simulated batting statistics for key players
	// Home team (Rajasthan Royals)
	info.BattingStats["Yashasvi Jaiswal"] = cricket.BattingStats{
		Runs:       42,
		Balls:      30,
		StrikeRate: 140.0,
		Boundaries: 5,
		Sixes:      2,
	}

	info.BattingStats["Jos Buttler"] = cricket.BattingStats{
		Runs:       26,
		Balls:      22,
		StrikeRate: 118.2,
		Boundaries: 3,
		Sixes:      1,
	}

	info.BattingStats["Sanju Samson"] = cricket.BattingStats{
		Runs:       19,
		Balls:      15,
		StrikeRate: 126.7,
		Boundaries: 2,
		Sixes:      1,
	}

	// Away team (Mumbai Indians) - higher scores as they won
	info.BattingStats["Rohit Sharma"] = cricket.BattingStats{
		Runs:       67,
		Balls:      42,
		StrikeRate: 159.5,
		Boundaries: 6,
		Sixes:      4,
	}

	info.BattingStats["Ishan Kishan"] = cricket.BattingStats{
		Runs:       54,
		Balls:      38,
		StrikeRate: 142.1,
		Boundaries: 5,
		Sixes:      3,
	}

	info.BattingStats["Suryakumar Yadav"] = cricket.BattingStats{
		Runs:       36,
		Balls:      24,
		StrikeRate: 150.0,
		Boundaries: 4,
		Sixes:      2,
	}

	// Simulated bowling statistics
	// Home team (Rajasthan Royals)
	info.BowlingStats["Trent Boult"] = cricket.BowlingStats{
		Overs:        4.0,
		RunsConceded: 38,
		Wickets:      2,
		Economy:      9.5,
	}

	info.BowlingStats["Yuzvendra Chahal"] = cricket.BowlingStats{
		Overs:        4.0,
		RunsConceded: 42,
		Wickets:      1,
		Economy:      10.5,
	}

	// Away team (Mumbai Indians)
	info.BowlingStats["Jasprit Bumrah"] = cricket.BowlingStats{
		Overs:        4.0,
		RunsConceded: 24,
		Wickets:      3,
		Economy:      6.0,
	}

	info.BowlingStats["Piyush Chawla"] = cricket.BowlingStats{
		Overs:        4.0,
		RunsConceded: 32,
		Wickets:      2,
		Economy:      8.0,
	}
}

// createMatchWinnerSelection creates and evaluates a match winner selection
func CreateMatchWinnerSelection(data cricket.CricketPrematchData, matchInfo cricket.DetailedMatchInfo) cricket.BetSelection {
	selection := cricket.BetSelection{
		Market:            "To Win the Match",
		MarketDescription: "Bet on which team will win the match",
		AvailableOptions:  []string{},
		ConfidenceLevel:   "High",
	}

	// Get all available options and odds for this market
	if len(data.Results) > 0 && len(data.Results[0].Main.SP.ToWinTheMatch.Odds) > 0 {
		// Get all available options
		for _, odd := range data.Results[0].Main.SP.ToWinTheMatch.Odds {
			odds, _ := strconv.ParseFloat(odd.Odds, 64)

			// Convert option names (1=Home, 2=Away) to team names
			optionName := ""
			if odd.Name == "1" {
				optionName = matchInfo.HomeTeam + " to win"
			} else if odd.Name == "2" {
				optionName = matchInfo.AwayTeam + " to win"
			} else {
				optionName = "Draw/Tie"
			}

			selection.AvailableOptions = append(selection.AvailableOptions,
				fmt.Sprintf("%s @ %.2f", optionName, odds))
		}

		// Our bet selection - let's pick the away team (Mumbai Indians)
		for _, odd := range data.Results[0].Main.SP.ToWinTheMatch.Odds {
			if odd.Name == "2" { // Away team selection
				odds, _ := strconv.ParseFloat(odd.Odds, 64)
				selection.Selection = matchInfo.AwayTeam + " to win"
				selection.Odds = odds
				selection.OddsDecimal = odds

				// Convert to American odds
				if odds >= 2.0 {
					americanOdds := (odds - 1) * 100
					selection.OddsAmerican = fmt.Sprintf("+%.0f", americanOdds)
				} else {
					americanOdds := -100 / (odds - 1)
					selection.OddsAmerican = fmt.Sprintf("%.0f", americanOdds)
				}

				// Convert to fractional odds (simplified)
				num := int((odds - 1) * 100)
				den := 100
				gcd := GreatestCommonDivisor(num, den)
				selection.OddsFractional = fmt.Sprintf("%d/%d", num/gcd, den/gcd)

				break
			}
		}
	}

	// Calculate potential profit with $100 stake
	selection.PotentialProfit = 100.0 * (selection.Odds - 1)

	// Set risk assessment based on odds
	if selection.Odds < 1.5 {
		selection.RiskAssessment = "Low Risk"
	} else if selection.Odds < 3.0 {
		selection.RiskAssessment = "Medium Risk"
	} else {
		selection.RiskAssessment = "High Risk"
	}

	// In cricket, the higher score wins
	if matchInfo.AwayScore > matchInfo.HomeScore {
		selection.IsWinner = true
		selection.Evaluation = fmt.Sprintf("%s won with score %d vs %d (margin: %d runs)",
			matchInfo.AwayTeam, matchInfo.AwayScore, matchInfo.HomeScore,
			matchInfo.AwayScore-matchInfo.HomeScore)
	} else if matchInfo.HomeScore > matchInfo.AwayScore {
		selection.IsWinner = false
		selection.Evaluation = fmt.Sprintf("%s won with score %d vs %d (margin: %d runs)",
			matchInfo.HomeTeam, matchInfo.HomeScore, matchInfo.AwayScore,
			matchInfo.HomeScore-matchInfo.AwayScore)
	} else {
		selection.IsWinner = false
		selection.Evaluation = "Match ended in a tie"
	}

	return selection
}

// createFirstOverRunsSelection creates and evaluates a first over total runs selection
func CreateFirstOverRunsSelection(data cricket.CricketPrematchData, matchInfo cricket.DetailedMatchInfo) cricket.BetSelection {
	selection := cricket.BetSelection{
		Market:            "1st Over Total Runs",
		MarketDescription: "Bet on the total number of runs scored in the first over",
		AvailableOptions:  []string{},
		ConfidenceLevel:   "Medium",
	}

	// Get all available options for this market
	if len(data.Results) > 0 && len(data.Results[0].FirstOver.SP.FirstOverTotalRuns.Odds) > 0 {
		for _, odd := range data.Results[0].FirstOver.SP.FirstOverTotalRuns.Odds {
			odds, _ := strconv.ParseFloat(odd.Odds, 64)
			selection.AvailableOptions = append(selection.AvailableOptions,
				fmt.Sprintf("%s %s @ %.2f", odd.Header, odd.Name, odds))
		}

		// For this example, we'll pick "Over 6.5" runs in the first over
		for _, odd := range data.Results[0].FirstOver.SP.FirstOverTotalRuns.Odds {
			if odd.Header == "Over" && odd.Name == "6.5" {
				odds, _ := strconv.ParseFloat(odd.Odds, 64)
				selection.Selection = fmt.Sprintf("Over %s runs in first over", odd.Name)
				selection.Odds = odds
				selection.OddsDecimal = odds

				// Convert to American odds
				if odds >= 2.0 {
					americanOdds := (odds - 1) * 100
					selection.OddsAmerican = fmt.Sprintf("+%.0f", americanOdds)
				} else {
					americanOdds := -100 / (odds - 1)
					selection.OddsAmerican = fmt.Sprintf("%.0f", americanOdds)
				}

				// Convert to fractional odds (simplified)
				num := int((odds - 1) * 100)
				den := 100
				gcd := GreatestCommonDivisor(num, den)
				selection.OddsFractional = fmt.Sprintf("%d/%d", num/gcd, den/gcd)

				break
			}
		}
	}

	// Calculate potential profit with $100 stake
	selection.PotentialProfit = 100.0 * (selection.Odds - 1)

	// Set risk assessment based on odds
	if selection.Odds < 1.5 {
		selection.RiskAssessment = "Low Risk"
	} else if selection.Odds < 3.0 {
		selection.RiskAssessment = "Medium Risk"
	} else {
		selection.RiskAssessment = "High Risk"
	}

	// For demonstration purposes, let's simulate the first over runs
	// In a real implementation, you'd need more detailed ball-by-ball data
	firstOverRuns := 8
	overUnderValue := 6.5

	if float64(firstOverRuns) > overUnderValue {
		selection.IsWinner = true
		selection.Evaluation = fmt.Sprintf("First over had %d runs (> %.1f). Breakdown: 1, 4, 0, 1, 2, 0",
			firstOverRuns, overUnderValue)
	} else {
		selection.IsWinner = false
		selection.Evaluation = fmt.Sprintf("First over had %d runs (<= %.1f)", firstOverRuns, overUnderValue)
	}

	return selection
}

// createFirstInningsScoreSelection creates and evaluates a first innings score selection
func CreateFirstInningsScoreSelection(data cricket.CricketPrematchData, matchInfo cricket.DetailedMatchInfo) cricket.BetSelection {
	selection := cricket.BetSelection{
		Market:            "1st Innings Score",
		MarketDescription: "Bet on whether the first innings score will be over or under a specific value",
		AvailableOptions:  []string{},
		ConfidenceLevel:   "Medium",
	}

	// Get all available options for this market
	if len(data.Results) > 0 && len(data.Results[0].Innings1.SP.FirstInningsScore.Odds) > 0 {
		for _, odd := range data.Results[0].Innings1.SP.FirstInningsScore.Odds {
			odds, _ := strconv.ParseFloat(odd.Odds, 64)
			selection.AvailableOptions = append(selection.AvailableOptions,
				fmt.Sprintf("%s %s @ %.2f", odd.Header, odd.Name, odds))
		}

		// For this example, we'll pick "Under 170.5" for the first innings score
		for _, odd := range data.Results[0].Innings1.SP.FirstInningsScore.Odds {
			if odd.Header == "Under" && odd.Name == "170.5" {
				odds, _ := strconv.ParseFloat(odd.Odds, 64)
				selection.Selection = fmt.Sprintf("Under %s runs in first innings", odd.Name)
				selection.Odds = odds
				selection.OddsDecimal = odds

				// Convert to American odds
				if odds >= 2.0 {
					americanOdds := (odds - 1) * 100
					selection.OddsAmerican = fmt.Sprintf("+%.0f", americanOdds)
				} else {
					americanOdds := -100 / (odds - 1)
					selection.OddsAmerican = fmt.Sprintf("%.0f", americanOdds)
				}

				// Convert to fractional odds (simplified)
				num := int((odds - 1) * 100)
				den := 100
				gcd := GreatestCommonDivisor(num, den)
				selection.OddsFractional = fmt.Sprintf("%d/%d", num/gcd, den/gcd)

				break
			}
		}
	}

	// Calculate potential profit with $100 stake
	selection.PotentialProfit = 100.0 * (selection.Odds - 1)

	// Set risk assessment based on odds
	if selection.Odds < 1.5 {
		selection.RiskAssessment = "Low Risk"
	} else if selection.Odds < 3.0 {
		selection.RiskAssessment = "Medium Risk"
	} else {
		selection.RiskAssessment = "High Risk"
	}

	// For demonstration, let's say the first innings score was 168
	firstInningsScore := 168
	thresholdValue := 170.5

	if float64(firstInningsScore) < thresholdValue {
		selection.IsWinner = true
		selection.Evaluation = fmt.Sprintf("First innings score was %d (< %.1f)",
			firstInningsScore, thresholdValue)
	} else {
		selection.IsWinner = false
		selection.Evaluation = fmt.Sprintf("First innings score was %d (>= %.1f)",
			firstInningsScore, thresholdValue)
	}

	return selection
}

// createFiftyToBeScored creates and evaluates whether a fifty will be scored
func CreateFiftyToBeScored(data cricket.CricketPrematchData, matchInfo cricket.DetailedMatchInfo) cricket.BetSelection {
	selection := cricket.BetSelection{
		Market:            "A Fifty to be Scored",
		MarketDescription: "Bet on whether any player will score fifty or more runs in the match",
		AvailableOptions:  []string{},
		ConfidenceLevel:   "High",
	}

	// Get all available options for this market
	if len(data.Results) > 0 && len(data.Results[0].Match.SP.AFiftyToBeScored.Odds) > 0 {
		for _, odd := range data.Results[0].Match.SP.AFiftyToBeScored.Odds {
			odds, _ := strconv.ParseFloat(odd.Odds, 64)
			selection.AvailableOptions = append(selection.AvailableOptions,
				fmt.Sprintf("%s @ %.2f", odd.Name, odds))
		}

		// We'll select "Yes" for a fifty to be scored
		for _, odd := range data.Results[0].Match.SP.AFiftyToBeScored.Odds {
			if odd.Name == "Yes" {
				odds, _ := strconv.ParseFloat(odd.Odds, 64)
				selection.Selection = "Yes - A fifty will be scored"
				selection.Odds = odds
				selection.OddsDecimal = odds

				// Convert to American odds
				if odds >= 2.0 {
					americanOdds := (odds - 1) * 100
					selection.OddsAmerican = fmt.Sprintf("+%.0f", americanOdds)
				} else {
					americanOdds := -100 / (odds - 1)
					selection.OddsAmerican = fmt.Sprintf("%.0f", americanOdds)
				}

				// Convert to fractional odds (simplified)
				num := int((odds - 1) * 100)
				den := 100
				gcd := GreatestCommonDivisor(num, den)
				selection.OddsFractional = fmt.Sprintf("%d/%d", num/gcd, den/gcd)

				break
			}
		}
	}

	// Calculate potential profit with $100 stake
	selection.PotentialProfit = 100.0 * (selection.Odds - 1)

	// Set risk assessment based on odds
	if selection.Odds < 1.5 {
		selection.RiskAssessment = "Low Risk"
	} else if selection.Odds < 3.0 {
		selection.RiskAssessment = "Medium Risk"
	} else {
		selection.RiskAssessment = "High Risk"
	}

	// Check if any player scored a fifty
	fiftyScored := false
	var fiftyScorers []string

	for player, stats := range matchInfo.BattingStats {
		if stats.Runs >= 50 {
			fiftyScored = true
			fiftyScorers = append(fiftyScorers, fmt.Sprintf("%s (%d)", player, stats.Runs))
		}
	}

	if fiftyScored {
		selection.IsWinner = true
		selection.Evaluation = fmt.Sprintf("A fifty was scored. Players with 50+ runs: %s",
			strings.Join(fiftyScorers, ", "))
	} else {
		selection.IsWinner = false
		selection.Evaluation = "No player scored fifty or more runs in the match"
	}

	return selection
}

// createSuperOverSelection creates and evaluates a super over selection
func CreateSuperOverSelection(data cricket.CricketPrematchData, matchInfo cricket.DetailedMatchInfo) cricket.BetSelection {
	selection := cricket.BetSelection{
		Market:            "To Go to Super Over",
		MarketDescription: "Bet on whether the match will go to a super over",
		AvailableOptions:  []string{},
		ConfidenceLevel:   "Low",
	}

	// Get all available options for this market
	if len(data.Results) > 0 && len(data.Results[0].Match.SP.ToGoToSuperOver.Odds) > 0 {
		for _, odd := range data.Results[0].Match.SP.ToGoToSuperOver.Odds {
			odds, _ := strconv.ParseFloat(odd.Odds, 64)
			selection.AvailableOptions = append(selection.AvailableOptions,
				fmt.Sprintf("%s @ %.2f", odd.Name, odds))
		}

		// We'll select "Yes" for the match to go to a super over
		for _, odd := range data.Results[0].Match.SP.ToGoToSuperOver.Odds {
			if odd.Name == "Yes" {
				odds, _ := strconv.ParseFloat(odd.Odds, 64)
				selection.Selection = "Yes - Match will go to super over"
				selection.Odds = odds
				selection.OddsDecimal = odds

				// Convert to American odds
				if odds >= 2.0 {
					americanOdds := (odds - 1) * 100
					selection.OddsAmerican = fmt.Sprintf("+%.0f", americanOdds)
				} else {
					americanOdds := -100 / (odds - 1)
					selection.OddsAmerican = fmt.Sprintf("%.0f", americanOdds)
				}

				// Convert to fractional odds (simplified)
				num := int((odds - 1) * 100)
				den := 100
				gcd := GreatestCommonDivisor(num, den)
				selection.OddsFractional = fmt.Sprintf("%d/%d", num/gcd, den/gcd)

				break
			}
		}
	}

	// Calculate potential profit with $100 stake
	selection.PotentialProfit = 100.0 * (selection.Odds - 1)

	// Set risk assessment based on odds
	if selection.Odds < 1.5 {
		selection.RiskAssessment = "Low Risk"
	} else if selection.Odds < 3.0 {
		selection.RiskAssessment = "Medium Risk"
	} else {
		selection.RiskAssessment = "High Risk"
	}

	// For this example, we'll say the match did not go to a super over
	superOver := false

	if superOver {
		selection.IsWinner = true
		selection.Evaluation = "The match went to a super over as the scores were tied"
	} else {
		selection.IsWinner = false
		selection.Evaluation = fmt.Sprintf("The match did not go to a super over. %s won by %d runs",
			matchInfo.AwayTeam, matchInfo.AwayScore-matchInfo.HomeScore)
	}

	return selection
}

// createMostSixesSelection creates and evaluates a most match sixes selection
func CreateMostSixesSelection(data cricket.CricketPrematchData, matchInfo cricket.DetailedMatchInfo) cricket.BetSelection {
	selection := cricket.BetSelection{
		Market:            "Most Match Sixes",
		MarketDescription: "Bet on which team will hit the most sixes in the match",
		AvailableOptions:  []string{},
		ConfidenceLevel:   "Medium",
	}

	// Get all available options for this market
	if len(data.Results) > 0 && len(data.Results[0].Match.SP.MostMatchSixes.Odds) > 0 {
		for _, odd := range data.Results[0].Match.SP.MostMatchSixes.Odds {
			odds, _ := strconv.ParseFloat(odd.Odds, 64)

			optionName := ""
			if odd.Name == "1" {
				optionName = matchInfo.HomeTeam
			} else if odd.Name == "2" {
				optionName = matchInfo.AwayTeam
			} else {
				optionName = "Tie"
			}

			selection.AvailableOptions = append(selection.AvailableOptions,
				fmt.Sprintf("%s @ %.2f", optionName, odds))
		}

		// We'll select the away team (Mumbai Indians) to hit the most sixes
		for _, odd := range data.Results[0].Match.SP.MostMatchSixes.Odds {
			if odd.Name == "2" { // Away team
				odds, _ := strconv.ParseFloat(odd.Odds, 64)
				selection.Selection = fmt.Sprintf("%s to hit most sixes", matchInfo.AwayTeam)
				selection.Odds = odds
				selection.OddsDecimal = odds

				// Convert to American odds
				if odds >= 2.0 {
					americanOdds := (odds - 1) * 100
					selection.OddsAmerican = fmt.Sprintf("+%.0f", americanOdds)
				} else {
					americanOdds := -100 / (odds - 1)
					selection.OddsAmerican = fmt.Sprintf("%.0f", americanOdds)
				}

				// Convert to fractional odds (simplified)
				num := int((odds - 1) * 100)
				den := 100
				gcd := GreatestCommonDivisor(num, den)
				selection.OddsFractional = fmt.Sprintf("%d/%d", num/gcd, den/gcd)

				break
			}
		}
	}

	// Calculate potential profit with $100 stake
	selection.PotentialProfit = 100.0 * (selection.Odds - 1)

	// Set risk assessment based on odds
	if selection.Odds < 1.5 {
		selection.RiskAssessment = "Low Risk"
	} else if selection.Odds < 3.0 {
		selection.RiskAssessment = "Medium Risk"
	} else {
		selection.RiskAssessment = "High Risk"
	}

	// Count sixes for each team
	homeSixes := 0
	awaySixes := 0

	for player, stats := range matchInfo.BattingStats {
		if strings.Contains(player, "Yashasvi") || strings.Contains(player, "Buttler") || strings.Contains(player, "Samson") {
			// These are home team players
			homeSixes += stats.Sixes
		} else {
			// These are away team players
			awaySixes += stats.Sixes
		}
	}

	if awaySixes > homeSixes {
		selection.IsWinner = true
		selection.Evaluation = fmt.Sprintf("%s hit more sixes (%d) than %s (%d)",
			matchInfo.AwayTeam, awaySixes, matchInfo.HomeTeam, homeSixes)
	} else if homeSixes > awaySixes {
		selection.IsWinner = false
		selection.Evaluation = fmt.Sprintf("%s hit more sixes (%d) than %s (%d)",
			matchInfo.HomeTeam, homeSixes, matchInfo.AwayTeam, awaySixes)
	} else {
		selection.IsWinner = false
		selection.Evaluation = fmt.Sprintf("Both teams hit the same number of sixes (%d)", homeSixes)
	}

	return selection
}

// createMostFoursSelection creates and evaluates a most match fours selection
func CreateMostFoursSelection(data cricket.CricketPrematchData, matchInfo cricket.DetailedMatchInfo) cricket.BetSelection {
	selection := cricket.BetSelection{
		Market:            "Most Match Fours",
		MarketDescription: "Bet on which team will hit the most fours in the match",
		AvailableOptions:  []string{},
		ConfidenceLevel:   "Medium",
	}

	// Get all available options for this market
	if len(data.Results) > 0 && len(data.Results[0].Match.SP.MostMatchFours.Odds) > 0 {
		for _, odd := range data.Results[0].Match.SP.MostMatchFours.Odds {
			odds, _ := strconv.ParseFloat(odd.Odds, 64)

			optionName := ""
			if odd.Name == "1" {
				optionName = matchInfo.HomeTeam
			} else if odd.Name == "2" {
				optionName = matchInfo.AwayTeam
			} else {
				optionName = "Tie"
			}

			selection.AvailableOptions = append(selection.AvailableOptions,
				fmt.Sprintf("%s @ %.2f", optionName, odds))
		}

		// We'll select the away team (Mumbai Indians) to hit the most fours
		for _, odd := range data.Results[0].Match.SP.MostMatchFours.Odds {
			if odd.Name == "2" { // Away team
				odds, _ := strconv.ParseFloat(odd.Odds, 64)
				selection.Selection = fmt.Sprintf("%s to hit most fours", matchInfo.AwayTeam)
				selection.Odds = odds
				selection.OddsDecimal = odds

				// Convert to American odds
				if odds >= 2.0 {
					americanOdds := (odds - 1) * 100
					selection.OddsAmerican = fmt.Sprintf("+%.0f", americanOdds)
				} else {
					americanOdds := -100 / (odds - 1)
					selection.OddsAmerican = fmt.Sprintf("%.0f", americanOdds)
				}

				// Convert to fractional odds (simplified)
				num := int((odds - 1) * 100)
				den := 100
				gcd := GreatestCommonDivisor(num, den)
				selection.OddsFractional = fmt.Sprintf("%d/%d", num/gcd, den/gcd)

				break
			}
		}
	}

	// Calculate potential profit with $100 stake
	selection.PotentialProfit = 100.0 * (selection.Odds - 1)

	// Set risk assessment based on odds
	if selection.Odds < 1.5 {
		selection.RiskAssessment = "Low Risk"
	} else if selection.Odds < 3.0 {
		selection.RiskAssessment = "Medium Risk"
	} else {
		selection.RiskAssessment = "High Risk"
	}

	// Count fours for each team
	homeFours := 0
	awayFours := 0

	for player, stats := range matchInfo.BattingStats {
		if strings.Contains(player, "Yashasvi") || strings.Contains(player, "Buttler") || strings.Contains(player, "Samson") {
			// These are home team players
			homeFours += stats.Boundaries
		} else {
			// These are away team players
			awayFours += stats.Boundaries
		}
	}

	if awayFours > homeFours {
		selection.IsWinner = true
		selection.Evaluation = fmt.Sprintf("%s hit more fours (%d) than %s (%d)",
			matchInfo.AwayTeam, awayFours, matchInfo.HomeTeam, homeFours)
	} else if homeFours > awayFours {
		selection.IsWinner = false
		selection.Evaluation = fmt.Sprintf("%s hit more fours (%d) than %s (%d)",
			matchInfo.HomeTeam, homeFours, matchInfo.AwayTeam, awayFours)
	} else {
		selection.IsWinner = false
		selection.Evaluation = fmt.Sprintf("Both teams hit the same number of fours (%d)", homeFours)
	}

	return selection
}

// createHundredToBeScored creates and evaluates a hundred to be scored selection
func CreateHundredToBeScored(data cricket.CricketPrematchData, matchInfo cricket.DetailedMatchInfo) cricket.BetSelection {
	selection := cricket.BetSelection{
		Market:            "A Hundred to be Scored",
		MarketDescription: "Bet on whether any player will score a century in the match",
		AvailableOptions:  []string{},
		ConfidenceLevel:   "Low",
	}

	// Get all available options for this market
	if len(data.Results) > 0 && len(data.Results[0].Match.SP.AHundredToBeScored.Odds) > 0 {
		for _, odd := range data.Results[0].Match.SP.AHundredToBeScored.Odds {
			odds, _ := strconv.ParseFloat(odd.Odds, 64)
			selection.AvailableOptions = append(selection.AvailableOptions,
				fmt.Sprintf("%s @ %.2f", odd.Name, odds))
		}

		// We'll select "No" for a hundred to be scored
		for _, odd := range data.Results[0].Match.SP.AHundredToBeScored.Odds {
			if odd.Name == "No" {
				odds, _ := strconv.ParseFloat(odd.Odds, 64)
				selection.Selection = "No - A hundred will not be scored"
				selection.Odds = odds
				selection.OddsDecimal = odds

				// Convert to American odds
				if odds >= 2.0 {
					americanOdds := (odds - 1) * 100
					selection.OddsAmerican = fmt.Sprintf("+%.0f", americanOdds)
				} else {
					americanOdds := -100 / (odds - 1)
					selection.OddsAmerican = fmt.Sprintf("%.0f", americanOdds)
				}

				// Convert to fractional odds (simplified)
				num := int((odds - 1) * 100)
				den := 100
				gcd := GreatestCommonDivisor(num, den)
				selection.OddsFractional = fmt.Sprintf("%d/%d", num/gcd, den/gcd)

				break
			}
		}
	}

	// Calculate potential profit with $100 stake
	selection.PotentialProfit = 100.0 * (selection.Odds - 1)

	// Set risk assessment based on odds
	if selection.Odds < 1.5 {
		selection.RiskAssessment = "Low Risk"
	} else if selection.Odds < 3.0 {
		selection.RiskAssessment = "Medium Risk"
	} else {
		selection.RiskAssessment = "High Risk"
	}

	// Check if any player scored a hundred
	hundredScored := false
	var centuryScorers []string

	for player, stats := range matchInfo.BattingStats {
		if stats.Runs >= 100 {
			hundredScored = true
			centuryScorers = append(centuryScorers, fmt.Sprintf("%s (%d)", player, stats.Runs))
		}
	}

	if hundredScored {
		selection.IsWinner = false
		selection.Evaluation = fmt.Sprintf("A hundred was scored. Players with 100+ runs: %s",
			strings.Join(centuryScorers, ", "))
	} else {
		selection.IsWinner = true
		selection.Evaluation = "No player scored a hundred in the match. Highest score was by Rohit Sharma (67 runs)"
	}

	return selection
}

func PrintMatchHeader(info cricket.DetailedMatchInfo) {
	fmt.Println("===========================================================")
	fmt.Println("               CRICKET MATCH ANALYSIS                      ")
	fmt.Println("===========================================================")
	fmt.Printf("Match: %s vs %s\n", info.HomeTeam, info.AwayTeam)
	fmt.Printf("Date: %s\n", info.MatchDate)
	fmt.Printf("Venue: %s, %s, %s (Capacity: %s)\n", info.Stadium, info.City, info.Country, info.Capacity)
	fmt.Printf("Competition: %s\n", info.LeagueName)
	fmt.Printf("Final Score: %s %d - %d %s\n", info.HomeTeam, info.HomeScore, info.AwayScore, info.AwayTeam)
	fmt.Println("-----------------------------------------------------------")
}

// printDetailedMatchStats prints detailed statistics about the match
func PrintDetailedMatchStats(info cricket.DetailedMatchInfo) {
	fmt.Println("\n===== KEY BATTING PERFORMANCES =====")
	fmt.Printf("%-20s %-8s %-8s %-8s %-8s %-8s\n", "PLAYER", "RUNS", "BALLS", "SR", "4s", "6s")
	fmt.Println("-------------------------------------------------------")

	// Print batting stats
	for player, stats := range info.BattingStats {
		fmt.Printf("%-20s %-8d %-8d %-8.2f %-8d %-8d\n",
			player, stats.Runs, stats.Balls, stats.StrikeRate, stats.Boundaries, stats.Sixes)
	}

	fmt.Println("\n===== KEY BOWLING PERFORMANCES =====")
	fmt.Printf("%-20s %-8s %-8s %-8s %-8s\n", "PLAYER", "OVERS", "RUNS", "WICKETS", "ECON")
	fmt.Println("-------------------------------------------------------")

	// Print bowling stats
	for player, stats := range info.BowlingStats {
		fmt.Printf("%-20s %-8.1f %-8d %-8d %-8.2f\n",
			player, stats.Overs, stats.RunsConceded, stats.Wickets, stats.Economy)
	}

	fmt.Println("-----------------------------------------------------------")
}

// printBettingEvaluationHeader prints the header for the betting evaluation
func PrintBettingEvaluationHeader() {
	fmt.Println("\n===========================================================")
	fmt.Println("               BETTING SELECTIONS EVALUATION                ")
	fmt.Println("===========================================================")
}

// printBetSelectionDetails prints the details of a betting selection
func PrintBetSelectionDetails(index int, selection cricket.BetSelection) {
	fmt.Printf("\n%d. %s\n", index, selection.Market)
	fmt.Printf("   Description: %s\n", selection.MarketDescription)
	fmt.Printf("   Selection: %s @ %.2f (Decimal: %.2f, American: %s, Fractional: %s)\n",
		selection.Selection, selection.Odds, selection.OddsDecimal, selection.OddsAmerican, selection.OddsFractional)
	fmt.Printf("   Potential Profit ($100 stake): $%.2f\n", selection.PotentialProfit)
	fmt.Printf("   Risk Assessment: %s\n", selection.RiskAssessment)
	fmt.Printf("   Confidence Level: %s\n", selection.ConfidenceLevel)

	fmt.Printf("   Available Options: %s\n", strings.Join(selection.AvailableOptions, " | "))

	if selection.IsWinner {
		fmt.Printf("   Result: WIN - %s\n", selection.Evaluation)
	} else {
		fmt.Printf("   Result: LOSS - %s\n", selection.Evaluation)
	}
}

// printBettingEvaluationSummary prints the summary of the betting evaluation
func PrintBettingEvaluationSummary(wins, total int, stake, returns, profitLoss, roi float64) {
	fmt.Println("\n===========================================================")
	fmt.Println("                      OVERALL SUMMARY                      ")
	fmt.Println("===========================================================")
	fmt.Printf("Winning Bets: %d/%d (%.1f%%)\n", wins, total, float64(wins)/float64(total)*100)
	fmt.Printf("Total Stake: $%.2f\n", stake*float64(total))
	fmt.Printf("Total Returns: $%.2f\n", returns)
	fmt.Printf("Profit/Loss: $%.2f\n", profitLoss)
	fmt.Printf("ROI: %.2f%%\n", roi)
	fmt.Println("-----------------------------------------------------------")
}

// printBettingHistory prints the betting history
func PrintBettingHistory(history []cricket.BettingHistory) {
	fmt.Println("\n===========================================================")
	fmt.Println("                   HISTORICAL PERFORMANCE                  ")
	fmt.Println("===========================================================")
	fmt.Printf("%-25s %-10s %-10s %-10s %-10s\n",
		"MARKET", "WIN RATE", "AVG ODDS", "BETS", "P/L")
	fmt.Println("-----------------------------------------------------------")

	for _, record := range history {
		fmt.Printf("%-25s %-10.1f%% %-10.2f %-10d $%-10.2f\n",
			record.Market, record.WinPercentage, record.AvgOdds, record.TotalBets, record.ProfitLoss)
	}

	fmt.Println("-----------------------------------------------------------")
}

// GreatestCommonDivisor calculates GCD using Euclidean algorithm
func GreatestCommonDivisor(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Helper function to convert decimal odds to American format
func DecimalToAmerican(odds float64) string {
	if odds >= 2.0 {
		return fmt.Sprintf("+%.0f", (odds-1)*100)
	}
	return fmt.Sprintf("%.0f", -100/(odds-1))
}

// Helper function to convert decimal odds to fractional format
func DecimalToFractional(odds float64) string {
	numerator := int((odds - 1) * 100)
	denominator := 100
	gcd := GreatestCommonDivisor(numerator, denominator)
	return fmt.Sprintf("%d/%d", numerator/gcd, denominator/gcd)
}
