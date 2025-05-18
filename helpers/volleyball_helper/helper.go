package volleyball_helper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/yesetoda/bet365-evaluator-go/models/volleyball"
)


func LoadVolleyballPrematchData(filename string) (*volleyball.PrematchData, error) {
	fileData, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var data volleyball.PrematchData
	if err := json.Unmarshal(fileData, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

func LoadVolleyballResultData(filename string) (*volleyball.ResultData, error) {
	fileData, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var data volleyball.ResultData
	if err := json.Unmarshal(fileData, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

func CalculateMatchStatistics(resultData *volleyball.ResultData) *volleyball.MatchStatistics {
	// Make sure we have data to process
	if len(resultData.Results) == 0 {
		log.Println("No match results found")
		return nil
	}

	result := resultData.Results[0]
	stats := &volleyball.MatchStatistics{}

	// Calculate total points for each set
	stats.TotalSet1Points = calculateSetPoints(result.Scores.Set1)
	stats.TotalSet2Points = calculateSetPoints(result.Scores.Set2)
	stats.TotalSet3Points = calculateSetPoints(result.Scores.Set3)
	stats.TotalSet4Points = calculateSetPoints(result.Scores.Set4)
	stats.TotalSet5Points = calculateSetPoints(result.Scores.Set5)

	// Calculate total match points
	stats.TotalMatchPoints = stats.TotalSet1Points + stats.TotalSet2Points + stats.TotalSet3Points + stats.TotalSet4Points + stats.TotalSet5Points

	// Parse match winner from SS field (format: "home_sets-away_sets")
	matchScoreParts := strings.Split(result.SS, "-")
	if len(matchScoreParts) == 2 {
		homeSetWins, _ := strconv.Atoi(matchScoreParts[0])
		awaySetWins, _ := strconv.Atoi(matchScoreParts[1])
		stats.HomeSetWins = homeSetWins
		stats.AwaySetWins = awaySetWins

		if homeSetWins > awaySetWins {
			stats.MatchWinner = "1" // Home team won
		} else {
			stats.MatchWinner = "2" // Away team won
		}

		// Calculate total sets played
		stats.TotalSets = homeSetWins + awaySetWins
	}

	// Determine maximum sets from "bestofsets" field
	if maxSets, err := strconv.Atoi(result.Extra.BestOfSets); err == nil {
		stats.MaximumSets = maxSets
	} else {
		stats.MaximumSets = 5 // Default to 5 sets if not specified
	}

	// Determine winners of each set
	homeSet1Points, _ := strconv.Atoi(result.Scores.Set1.Home)
	awaySet1Points, _ := strconv.Atoi(result.Scores.Set1.Away)
	homeSet2Points, _ := strconv.Atoi(result.Scores.Set2.Home)
	awaySet2Points, _ := strconv.Atoi(result.Scores.Set2.Away)
	homeSet3Points, _ := strconv.Atoi(result.Scores.Set3.Home)
	awaySet3Points, _ := strconv.Atoi(result.Scores.Set3.Away)
	homeSet4Points, _ := strconv.Atoi(result.Scores.Set4.Home)
	awaySet4Points, _ := strconv.Atoi(result.Scores.Set4.Away)
	homeSet5Points, _ := strconv.Atoi(result.Scores.Set5.Home)
	awaySet5Points, _ := strconv.Atoi(result.Scores.Set5.Away)

	// Set 1 winner
	if homeSet1Points > awaySet1Points {
		stats.Set1Winner = "1" // Home team won set 1
	} else if awaySet1Points > homeSet1Points {
		stats.Set1Winner = "2" // Away team won set 1
	}

	// Set 2 winner
	if homeSet2Points > awaySet2Points {
		stats.Set2Winner = "1" // Home team won set 2
	} else if awaySet2Points > homeSet2Points {
		stats.Set2Winner = "2" // Away team won set 2
	}

	// Set 3 winner
	if homeSet3Points > awaySet3Points {
		stats.Set3Winner = "1" // Home team won set 3
	} else if awaySet3Points > homeSet3Points {
		stats.Set3Winner = "2" // Away team won set 3
	}

	// Set 4 winner
	if homeSet4Points > awaySet4Points {
		stats.Set4Winner = "1" // Home team won set 4
	} else if awaySet4Points > homeSet4Points {
		stats.Set4Winner = "2" // Away team won set 4
	}

	// Set 5 winner
	if homeSet5Points > awaySet5Points {
		stats.Set5Winner = "1" // Home team won set 5
	} else if awaySet5Points > homeSet5Points {
		stats.Set5Winner = "2" // Away team won set 5
	}

	// Calculate if sets went to extra points (in volleyball, extra points occur when score goes beyond standard winning score)
	// For standard sets (1-4), the winning score is typically 25 points
	// For the final set (5), the winning score is typically 15 points
	stats.Set1ExtraPoints = homeSet1Points > 25 || awaySet1Points > 25
	stats.Set2ExtraPoints = homeSet2Points > 25 || awaySet2Points > 25
	stats.Set3ExtraPoints = homeSet3Points > 25 || awaySet3Points > 25
	stats.Set4ExtraPoints = homeSet4Points > 25 || awaySet4Points > 25
	stats.Set5ExtraPoints = homeSet5Points > 15 || awaySet5Points > 15

	// Determine correct set score (format: "winner sets-loser sets")
	stats.CorrectSetScore = fmt.Sprintf("%s %d-%d", stats.MatchWinner, max(stats.HomeSetWins, stats.AwaySetWins), min(stats.HomeSetWins, stats.AwaySetWins))

	return stats
}

func CreateBetSelections(prematchData *volleyball.PrematchData, stakeAmount float64) []volleyball.BetSelection {
	selections := []volleyball.BetSelection{}

	// Make sure we have data to process
	if len(prematchData.Results) == 0 {
		log.Println("No prematch results found")
		return selections
	}

	result := prematchData.Results[0]

	// 1. Match Winner (1X2 equivalent in volleyball)
	for _, odds := range result.Main.Sp.GameLines.Odds {
		if odds.Header != "" && odds.Handicap == "" && odds.Odds != "" {
			// Skip parent odds entries with empty odds value
			if odds.Odds == "" {
				continue
			}
			oddsValue, err := strconv.ParseFloat(odds.Odds, 64)
			if err == nil && oddsValue > 0 {
				selections = append(selections, volleyball.BetSelection{
					Market:      "Match Winner",
					MarketID:    result.Main.Sp.GameLines.ID,
					Selection:   odds.Header,
					SelectionID: odds.ID,
					Odds:        oddsValue,
					Handicap:    "",
					StakeAmount: stakeAmount,
				})
			}
		}
	}

	// 2. Handicap
	for _, odds := range result.Main.Sp.GameLines.Odds {
		if odds.Header != "" && odds.Handicap != "" && (strings.Contains(odds.Handicap, "-") || strings.Contains(odds.Handicap, "+")) {
			oddsValue, err := strconv.ParseFloat(odds.Odds, 64)
			if err == nil && oddsValue > 0 {
				selections = append(selections, volleyball.BetSelection{
					Market:      "Handicap",
					MarketID:    result.Main.Sp.GameLines.ID,
					Selection:   odds.Header,
					SelectionID: odds.ID,
					Odds:        oddsValue,
					Handicap:    odds.Handicap,
					StakeAmount: stakeAmount,
				})
			}
		}
	}

	// 3. Total Points
	for _, odds := range result.Main.Sp.GameLines.Odds {
		if odds.Handicap != "" && (strings.HasPrefix(odds.Handicap, "O ") || strings.HasPrefix(odds.Handicap, "U ")) {
			oddsValue, err := strconv.ParseFloat(odds.Odds, 64)
			if err == nil && oddsValue > 0 {
				selections = append(selections, volleyball.BetSelection{
					Market:      "Total Points",
					MarketID:    result.Main.Sp.GameLines.ID,
					Selection:   odds.Handicap,
					SelectionID: odds.ID,
					Odds:        oddsValue,
					Handicap:    odds.Handicap,
					StakeAmount: stakeAmount,
				})
			}
		}
	}

	// 4. Correct Set Score
	for _, odds := range result.Main.Sp.CorrectSetScore.Odds {
		oddsValue, err := strconv.ParseFloat(odds.Odds, 64)
		if err == nil && oddsValue > 0 {
			selections = append(selections, volleyball.BetSelection{
				Market:      "Correct Set Score",
				MarketID:    result.Main.Sp.CorrectSetScore.ID,
				Selection:   fmt.Sprintf("%s %s", odds.Header, odds.Name),
				SelectionID: odds.ID,
				Odds:        oddsValue,
				Handicap:    "",
				StakeAmount: stakeAmount,
			})
		}
	}

	// 5. Set 1 Winner
	for _, other := range result.Others {
		if len(other.Sp.Set1Lines.Odds) > 0 {
			for _, odds := range other.Sp.Set1Lines.Odds {
				if odds.Header != "" && odds.Handicap == "" && odds.Name == "Winner" {
					oddsValue, err := strconv.ParseFloat(odds.Odds, 64)
					if err == nil && oddsValue > 0 {
						selections = append(selections, volleyball.BetSelection{
							Market:      "Set 1 Winner",
							MarketID:    other.Sp.Set1Lines.ID,
							Selection:   odds.Header,
							SelectionID: odds.ID,
							Odds:        oddsValue,
							Handicap:    "",
							StakeAmount: stakeAmount,
						})
					}
				}
			}
		}
	}

	// 6. Set 1 Total Points
	for _, other := range result.Others {
		if len(other.Sp.Set1Lines.Odds) > 0 {
			for _, odds := range other.Sp.Set1Lines.Odds {
				if odds.Handicap != "" && (strings.HasPrefix(odds.Handicap, "O ") || strings.HasPrefix(odds.Handicap, "U ")) {
					oddsValue, err := strconv.ParseFloat(odds.Odds, 64)
					if err == nil && oddsValue > 0 {
						selections = append(selections, volleyball.BetSelection{
							Market:      "Set 1 Total Points",
							MarketID:    other.Sp.Set1Lines.ID,
							Selection:   odds.Handicap,
							SelectionID: odds.ID,
							Odds:        oddsValue,
							Handicap:    odds.Handicap,
							StakeAmount: stakeAmount,
						})
					}
				}
			}
		}
	}

	// 7. Match Total Odd/Even
	for _, other := range result.Others {
		if len(other.Sp.MatchTotalOddEven.Odds) > 0 {
			for _, odds := range other.Sp.MatchTotalOddEven.Odds {
				oddsValue, err := strconv.ParseFloat(odds.Odds, 64)
				if err == nil && oddsValue > 0 {
					selections = append(selections, volleyball.BetSelection{
						Market:      "Match Total Odd/Even",
						MarketID:    other.Sp.MatchTotalOddEven.ID,
						Selection:   odds.Name,
						SelectionID: odds.ID,
						Odds:        oddsValue,
						Handicap:    "",
						StakeAmount: stakeAmount,
					})
				}
			}
		}
	}

	// 8. Set 1 Extra Points
	for _, other := range result.Others {
		if len(other.Sp.Set1ToGoToExtraPoints.Odds) > 0 {
			for _, odds := range other.Sp.Set1ToGoToExtraPoints.Odds {
				oddsValue, err := strconv.ParseFloat(odds.Odds, 64)
				if err == nil && oddsValue > 0 {
					selections = append(selections, volleyball.BetSelection{
						Market:      "Set 1 Extra Points",
						MarketID:    other.Sp.Set1ToGoToExtraPoints.ID,
						Selection:   odds.Name,
						SelectionID: odds.ID,
						Odds:        oddsValue,
						Handicap:    "",
						StakeAmount: stakeAmount,
					})
				}
			}
		}
	}

	// 9. Set 1 Total Odd/Even
	for _, other := range result.Others {
		if len(other.Sp.Set1TotalOddEven.Odds) > 0 {
			for _, odds := range other.Sp.Set1TotalOddEven.Odds {
				oddsValue, err := strconv.ParseFloat(odds.Odds, 64)
				if err == nil && oddsValue > 0 {
					selections = append(selections, volleyball.BetSelection{
						Market:      "Set 1 Total Odd/Even",
						MarketID:    other.Sp.Set1TotalOddEven.ID,
						Selection:   odds.Name,
						SelectionID: odds.ID,
						Odds:        oddsValue,
						Handicap:    "",
						StakeAmount: stakeAmount,
					})
				}
			}
		}
	}

	// 10. Double Chance (custom implementation for volleyball)
	// In volleyball, this would be interpreted as backing both teams to win
	// Let's create synthetic Double Chance markets
	homeWinOdds := 0.0
	awayWinOdds := 0.0

	for _, odds := range result.Main.Sp.GameLines.Odds {
		if odds.Header == "1" && odds.Handicap == "" {
			homeWinOdds, _ = strconv.ParseFloat(odds.Odds, 64)
		} else if odds.Header == "2" && odds.Handicap == "" {
			awayWinOdds, _ = strconv.ParseFloat(odds.Odds, 64)
		}
	}

	// Calculate Combined Probability for Double Chance
	if homeWinOdds > 0 && awayWinOdds > 0 {
		homeProbability := 1.0 / homeWinOdds
		awayProbability := 1.0 / awayWinOdds

		// Calculate Double Chance odds using combined probability
		// Formula: 1 / (P(Home) + P(Away))
		doubleChanceOdds := 1.0 / (homeProbability + awayProbability)

		selections = append(selections, volleyball.BetSelection{
			Market:      "Double Chance",
			MarketID:    "custom_double_chance",
			Selection:   "1-2", // Both teams to win
			SelectionID: "custom_dc_1",
			Odds:        float64(int(doubleChanceOdds*100)) / 100, // Round to 2 decimal places
			Handicap:    "",
			StakeAmount: stakeAmount,
		})
	}

	return selections
}

func EvaluateBetSelections(selections []volleyball.BetSelection, resultData *volleyball.ResultData, matchStats *volleyball.MatchStatistics) []volleyball.EvaluationResult {
	evaluations := []volleyball.EvaluationResult{}

	// Make sure we have data to process
	if matchStats == nil || len(resultData.Results) == 0 {
		log.Println("No match results or statistics available")
		return evaluations
	}

	result := resultData.Results[0]

	// Evaluate each selection
	for _, selection := range selections {
		evaluation := volleyball.EvaluationResult{
			BetSelection:       selection,
			IsWin:              false,
			ImpliedProbability: 1.0 / selection.Odds * 100, // Calculate implied probability
		}

		switch selection.Market {
		case "Match Winner":
			evaluation.IsWin = selection.Selection == matchStats.MatchWinner
			evaluation.Explanation = fmt.Sprintf("Match result: %s-%s. Winner: Team %s. User bet: Team %s to win.",
				result.Scores.Set1.Home, result.Scores.Set1.Away, matchStats.MatchWinner, selection.Selection)

		case "Handicap":
			// Parse handicap value
			handicapValue := 0.0
			if strings.Contains(selection.Handicap, "-") {
				parts := strings.Split(selection.Handicap, "-")
				if len(parts) > 1 {
					handicapValue, _ = strconv.ParseFloat(parts[1], 64)
					handicapValue = -handicapValue
				}
			} else if strings.Contains(selection.Handicap, "+") {
				parts := strings.Split(selection.Handicap, "+")
				if len(parts) > 1 {
					handicapValue, _ = strconv.ParseFloat(parts[1], 64)
				}
			}

			// Apply handicap to set difference
			setDiff := matchStats.HomeSetWins - matchStats.AwaySetWins
			adjustedDiff := float64(setDiff)

			if selection.Selection == "1" { // Home team
				adjustedDiff += handicapValue
				evaluation.IsWin = adjustedDiff > 0
			} else if selection.Selection == "2" { // Away team
				adjustedDiff = -adjustedDiff + handicapValue
				evaluation.IsWin = adjustedDiff > 0
			}

			homeTeam := result.Home.Name
			awayTeam := result.Away.Name
			teamName := homeTeam
			if selection.Selection == "2" {
				teamName = awayTeam
			}

			evaluation.Explanation = fmt.Sprintf("Match result: %d-%d sets. Actual set difference: %d. Applied handicap %s to %s: adjusted difference %.1f. User bet: %s with handicap %s.",
				matchStats.HomeSetWins, matchStats.AwaySetWins, setDiff, selection.Handicap, teamName, adjustedDiff, teamName, selection.Handicap)

		case "Total Points":
			// Parse total value
			totalValue := 0.0
			isOver := strings.HasPrefix(selection.Handicap, "O ")

			if isOver {
				parts := strings.Split(selection.Handicap, "O ")
				if len(parts) > 1 {
					totalValue, _ = strconv.ParseFloat(parts[1], 64)
				}
			} else {
				parts := strings.Split(selection.Handicap, "U ")
				if len(parts) > 1 {
					totalValue, _ = strconv.ParseFloat(parts[1], 64)
				}
			}

			if isOver {
				evaluation.IsWin = float64(matchStats.TotalMatchPoints) > totalValue
			} else {
				evaluation.IsWin = float64(matchStats.TotalMatchPoints) < totalValue
			}

			evaluation.Explanation = fmt.Sprintf("Total match points: %d. User bet: %s (threshold: %.1f). Result: %s",
				matchStats.TotalMatchPoints, selection.Selection, totalValue, getResultText(evaluation.IsWin))

		case "Correct Set Score":
			// Parse team and score from selection
			parts := strings.Split(selection.Selection, " ")
			if len(parts) < 2 {
				evaluation.Explanation = "Invalid selection format"
				break
			}

			team := parts[0]
			scoreParts := strings.Split(parts[1], "-")
			if len(scoreParts) < 2 {
				evaluation.Explanation = "Invalid score format"
				break
			}

			winSets, _ := strconv.Atoi(scoreParts[0])
			loseSets, _ := strconv.Atoi(scoreParts[1])

			// Check if the prediction matches the actual result
			if team == matchStats.MatchWinner &&
				winSets == max(matchStats.HomeSetWins, matchStats.AwaySetWins) &&
				loseSets == min(matchStats.HomeSetWins, matchStats.AwaySetWins) {
				evaluation.IsWin = true
			}

			teamName := "Home team"
			if team == "2" {
				teamName = "Away team"
			}

			// Get actual team names
			homeTeam := result.Home.Name
			awayTeam := result.Away.Name
			selectedTeam := homeTeam
			if team == "2" {
				selectedTeam = awayTeam
			}

			evaluation.Explanation = fmt.Sprintf("Match result: %d-%d. User bet: %s (%s) to win %s. Actual winner: %s (%s) with score %d-%d. Result: %s",
				matchStats.HomeSetWins, matchStats.AwaySetWins, teamName, selectedTeam, parts[1],
				getTeamType(matchStats.MatchWinner), getTeamName(matchStats.MatchWinner, homeTeam, awayTeam),
				max(matchStats.HomeSetWins, matchStats.AwaySetWins), min(matchStats.HomeSetWins, matchStats.AwaySetWins),
				getResultText(evaluation.IsWin))

		case "Set 1 Winner":
			evaluation.IsWin = selection.Selection == matchStats.Set1Winner

			homeTeam := result.Home.Name
			awayTeam := result.Away.Name
			teamName := getTeamName(selection.Selection, homeTeam, awayTeam)
			actualWinner := getTeamName(matchStats.Set1Winner, homeTeam, awayTeam)

			evaluation.Explanation = fmt.Sprintf("Set 1 result: %s-%s. Winner: %s. User bet: %s to win Set 1. Result: %s",
				result.Scores.Set1.Home, result.Scores.Set1.Away, actualWinner, teamName, getResultText(evaluation.IsWin))

		case "Set 1 Total Points":
			// Parse total value
			totalValue := 0.0
			isOver := strings.HasPrefix(selection.Handicap, "O ")

			if isOver {
				parts := strings.Split(selection.Handicap, "O ")
				if len(parts) > 1 {
					totalValue, _ = strconv.ParseFloat(parts[1], 64)
				}
			} else {
				parts := strings.Split(selection.Handicap, "U ")
				if len(parts) > 1 {
					totalValue, _ = strconv.ParseFloat(parts[1], 64)
				}
			}

			if isOver {
				evaluation.IsWin = float64(matchStats.TotalSet1Points) > totalValue
			} else {
				evaluation.IsWin = float64(matchStats.TotalSet1Points) < totalValue
			}

			evaluation.Explanation = fmt.Sprintf("Set 1 total points: %d. User bet: %s (threshold: %.1f). Result: %s",
				matchStats.TotalSet1Points, selection.Selection, totalValue, getResultText(evaluation.IsWin))

		case "Match Total Odd/Even":
			isOdd := matchStats.TotalMatchPoints%2 == 1
			evaluation.IsWin = (selection.Selection == "Odd" && isOdd) || (selection.Selection == "Even" && !isOdd)

			oddEvenText := "Even"
			if isOdd {
				oddEvenText = "Odd"
			}

			evaluation.Explanation = fmt.Sprintf("Total match points: %d (%s). User bet: %s. Result: %s",
				matchStats.TotalMatchPoints, oddEvenText, selection.Selection, getResultText(evaluation.IsWin))

		case "Set 1 Extra Points":
			evaluation.IsWin = (selection.Selection == "Yes" && matchStats.Set1ExtraPoints) ||
				(selection.Selection == "No" && !matchStats.Set1ExtraPoints)

			extraPointsText := "No"
			if matchStats.Set1ExtraPoints {
				extraPointsText = "Yes"
			}

			evaluation.Explanation = fmt.Sprintf("Set 1 had extra points: %s. User bet: %s. Result: %s",
				extraPointsText, selection.Selection, getResultText(evaluation.IsWin))

		case "Set 1 Total Odd/Even":
			isOdd := matchStats.TotalSet1Points%2 == 1
			evaluation.IsWin = (selection.Selection == "Odd" && isOdd) || (selection.Selection == "Even" && !isOdd)

			oddEvenText := "Even"
			if isOdd {
				oddEvenText = "Odd"
			}

			evaluation.Explanation = fmt.Sprintf("Set 1 total points: %d (%s). User bet: %s. Result: %s",
				matchStats.TotalSet1Points, oddEvenText, selection.Selection, getResultText(evaluation.IsWin))

		case "Double Chance":
			// For volleyball, Double Chance would mean backing both teams
			// In this implementation, Double Chance always loses because only one team can win
			evaluation.IsWin = false
			evaluation.Explanation = fmt.Sprintf("Double Chance is not applicable in volleyball as only one team can win. User bet: %s. Result: Loss",
				selection.Selection)
		}

		// Calculate profit/loss and return amount
		if evaluation.IsWin {
			evaluation.ProfitLoss = selection.StakeAmount * (selection.Odds - 1.0)
			evaluation.ReturnAmount = selection.StakeAmount * selection.Odds
		} else {
			evaluation.ProfitLoss = -selection.StakeAmount
			evaluation.ReturnAmount = 0.0
		}

		evaluations = append(evaluations, evaluation)
	}

	return evaluations
}

func DisplayResults(evaluations []volleyball.EvaluationResult, resultData *volleyball.ResultData, matchStats *volleyball.MatchStatistics) {
	// Make sure we have data to process
	if matchStats == nil || len(resultData.Results) == 0 || len(evaluations) == 0 {
		log.Println("No data to display")
		return
	}

	result := resultData.Results[0]

	fmt.Println("======================== MATCH SUMMARY ========================")
	fmt.Printf("Match: %s vs %s\n", result.Home.Name, result.Away.Name)
	fmt.Printf("League: %s\n", result.League.Name)
	fmt.Printf("Date: %s\n", formatTimestamp(result.Time))
	fmt.Printf("Final Score: %s\n", result.SS)
	fmt.Printf("\nSet scores:\n")
	fmt.Printf("  Set 1: %s-%s\n", result.Scores.Set1.Home, result.Scores.Set1.Away)
	fmt.Printf("  Set 2: %s-%s\n", result.Scores.Set2.Home, result.Scores.Set2.Away)
	fmt.Printf("  Set 3: %s-%s\n", result.Scores.Set3.Home, result.Scores.Set3.Away)

	if matchStats.TotalSets > 3 {
		fmt.Printf("  Set 4: %s-%s\n", result.Scores.Set4.Home, result.Scores.Set4.Away)
	}
	if matchStats.TotalSets > 4 {
		fmt.Printf("  Set 5: %s-%s\n", result.Scores.Set5.Home, result.Scores.Set5.Away)
	}

	// Display key statistics
	fmt.Println("\n======================== KEY STATISTICS ========================")
	fmt.Printf("Total Match Points: %d\n", matchStats.TotalMatchPoints)
	fmt.Printf("Set 1 Points: %d (%s)\n", matchStats.TotalSet1Points, getOddEvenText(matchStats.TotalSet1Points))
	fmt.Printf("Set 1 Extra Points: %t\n", matchStats.Set1ExtraPoints)
	fmt.Printf("Total Match Points Odd/Even: %s\n", getOddEvenText(matchStats.TotalMatchPoints))
	fmt.Printf("Match Winner: %s (%s)\n", getTeamType(matchStats.MatchWinner), getTeamName(matchStats.MatchWinner, result.Home.Name, result.Away.Name))
	fmt.Printf("Correct Set Score: %s\n", fmt.Sprintf("%d-%d", max(matchStats.HomeSetWins, matchStats.AwaySetWins), min(matchStats.HomeSetWins, matchStats.AwaySetWins)))

	// Display bet results
	fmt.Println("\n======================== BET RESULTS ========================")

	totalStake := 0.0
	totalProfit := 0.0
	winCount := 0

	for _, eval := range evaluations {
		resultText := "LOSS"
		if eval.IsWin {
			resultText = "WIN"
			winCount++
		}

		fmt.Printf("\n----- %s -----\n", eval.BetSelection.Market)
		fmt.Printf("Selection: %s @ %.2f\n", eval.BetSelection.Selection, eval.BetSelection.Odds)
		fmt.Printf("Stake: $%.2f\n", eval.BetSelection.StakeAmount)
		fmt.Printf("Result: %s\n", resultText)
		fmt.Printf("Profit/Loss: $%.2f\n", eval.ProfitLoss)
		fmt.Printf("Implied Probability: %.2f%%\n", eval.ImpliedProbability)
		fmt.Printf("Explanation: %s\n", eval.Explanation)

		totalStake += eval.BetSelection.StakeAmount
		totalProfit += eval.ProfitLoss
	}

	// Display summary statistics
	winRate := float64(winCount) / float64(len(evaluations)) * 100
	roi := totalProfit / totalStake * 100

	fmt.Println("\n===================== BETTING SUMMARY =====================")
	fmt.Printf("Total Bets: %d\n", len(evaluations))
	fmt.Printf("Winning Bets: %d (%.2f%%)\n", winCount, winRate)
	fmt.Printf("Total Stake: $%.2f\n", totalStake)
	fmt.Printf("Total Profit/Loss: $%.2f\n", totalProfit)
	fmt.Printf("ROI: %.2f%%\n", roi)
}

func calculateSetPoints(set volleyball.SetScore) int {
	homePoints, _ := strconv.Atoi(set.Home)
	awayPoints, _ := strconv.Atoi(set.Away)
	return homePoints + awayPoints
}

func getResultText(isWin bool) string {
	if isWin {
		return "WIN"
	}
	return "LOSS"
}

func getOddEvenText(value int) string {
	if value%2 == 0 {
		return "Even"
	}
	return "Odd"
}

func getTeamType(team string) string {
	if team == "1" {
		return "Home team"
	}
	return "Away team"
}

func getTeamName(team string, homeName string, awayName string) string {
	if team == "1" {
		return homeName
	}
	return awayName
}

func formatTimestamp(timestamp string) string {
	// Convert unix timestamp to Go time
	i, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return timestamp
	}

	t := time.Unix(i, 0)
	return t.Format("January 2, 2006 15:04:05")
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
