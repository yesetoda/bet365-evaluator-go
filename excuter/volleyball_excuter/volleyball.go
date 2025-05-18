package volleyball_excuter
import (
	"log"
	"os"

	"github.com/yesetoda/bet365-evaluator-go/helpers/volleyball_helper"
)

func VolleyballExecutor() {
	// Check if file paths are provided as command-line arguments
	var prematchFilePath, resultFilePath string

	if len(os.Args) > 2 {
		prematchFilePath = os.Args[1]
		resultFilePath = os.Args[2]
	} else {
		// Default file paths
		prematchFilePath = "data/volleyball_prematch.json"
		resultFilePath = "data/volleyball_result.json"
	}

	// Load prematch data
	prematchData, err := volleyball_helper.LoadVolleyballPrematchData(prematchFilePath)
	if err != nil {
		log.Fatalf("Failed to load prematch data: %v", err)
	}

	// Load result data
	resultData, err := volleyball_helper.LoadVolleyballResultData(resultFilePath)
	if err != nil {
		log.Fatalf("Failed to load result data: %v", err)
	}

	// Simulate stake amount for each bet
	stakeAmount := 100.0 // Default stake amount of $100

	// Create bet selections
	selections := volleyball_helper.CreateBetSelections(prematchData, stakeAmount)

	// Calculate match statistics
	matchStats := volleyball_helper.CalculateMatchStatistics(resultData)

	// Evaluate bet selections
	evaluations := volleyball_helper.EvaluateBetSelections(selections, resultData, matchStats)

	// Display results
	volleyball_helper.DisplayResults(evaluations, resultData, matchStats)
}
