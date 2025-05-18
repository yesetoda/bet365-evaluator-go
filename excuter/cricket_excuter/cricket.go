package cricket_excuter

import (
	"log"
	"math/rand"
	"time"

	"github.com/yesetoda/bet365-evaluator-go/helpers/cricket_helper"
	"github.com/yesetoda/bet365-evaluator-go/models/cricket"
)

func CricketExecutor() {
	// Set up logging to file

	// Seed the random number generator for simulated data
	rand.Seed(time.Now().UnixNano())

	// Load the JSON data from files
	resultData, err := cricket_helper.LoadCricketResultData("data/cricket_result.json")
	if err != nil {
		log.Fatalf("Failed to load result data: %v", err)
	}

	prematchData, err := cricket_helper.LoadCricketPrematchData("data/cricket_prematch.json")
	if err != nil {
		log.Fatalf("Failed to load prematch data: %v", err)
	}

	// Extract and process the match information
	matchInfo := cricket_helper.ExtractDetailedMatchInfo(resultData)

	// Print detailed match information
	cricket_helper.PrintMatchHeader(matchInfo)

	// Generate simulated statistics for demonstration
	cricket_helper.SimulateDetailedMatchStats(&matchInfo)

	// Print extended match statistics
	cricket_helper.PrintDetailedMatchStats(matchInfo)

	// Create a list of bet selections and evaluate them
	betSelections := []cricket.BetSelection{}

	// Market 1: Match Winner (To Win the Match)
	winnerSelection := cricket_helper.CreateMatchWinnerSelection(prematchData, matchInfo)
	betSelections = append(betSelections, winnerSelection)

	// Market 2: First Over Total Runs
	firstOverRunsSelection := cricket_helper.CreateFirstOverRunsSelection(prematchData, matchInfo)
	betSelections = append(betSelections, firstOverRunsSelection)

	// Market 3: 1st Innings Score
	inningsScoreSelection := cricket_helper.CreateFirstInningsScoreSelection(prematchData, matchInfo)
	betSelections = append(betSelections, inningsScoreSelection)

	// Market 4: A Fifty to be Scored
	fiftySelection := cricket_helper.CreateFiftyToBeScored(prematchData, matchInfo)
	betSelections = append(betSelections, fiftySelection)

	// Market 5: Super Over
	superOverSelection := cricket_helper.CreateSuperOverSelection(prematchData, matchInfo)
	betSelections = append(betSelections, superOverSelection)

	// Market 6: Most Match Sixes
	mostSixesSelection := cricket_helper.CreateMostSixesSelection(prematchData, matchInfo)
	betSelections = append(betSelections, mostSixesSelection)

	// Market 7: Most Match Fours
	mostFoursSelection := cricket_helper.CreateMostFoursSelection(prematchData, matchInfo)
	betSelections = append(betSelections, mostFoursSelection)

	// Market 8: A Hundred to be Scored
	hundredSelection := cricket_helper.CreateHundredToBeScored(prematchData, matchInfo)
	betSelections = append(betSelections, hundredSelection)

	// Display evaluation results
	cricket_helper.PrintBettingEvaluationHeader()

	for i, selection := range betSelections {
		cricket_helper.PrintBetSelectionDetails(i+1, selection)
	}

	// Overall summary and additional metrics
	wins := 0
	totalStake := 100.0 // Assuming equal stakes for simplicity
	totalReturns := 0.0

	for _, bet := range betSelections {
		if bet.IsWinner {
			wins++
			totalReturns += totalStake * bet.Odds
		}
	}

	profitLoss := totalReturns - (totalStake * float64(len(betSelections)))
	roi := (profitLoss / (totalStake * float64(len(betSelections)))) * 100

	cricket_helper.PrintBettingEvaluationSummary(wins, len(betSelections), totalStake, totalReturns, profitLoss, roi)

	log.Println("Completed cricket betting evaluation at", time.Now().Format(time.RFC1123))
}
