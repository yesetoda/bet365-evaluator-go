# Bet365 Market Analysis Engine ⚽🏆

[![Go Report Card](https://goreportcard.com/badge/github.com/yourusername/bet365-evaluator-go)](https://goreportcard.com/report/github.com/yourusername/bet365-evaluator-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A command-line application for analyzing Bet365 sports betting markets, demonstrating advanced JSON parsing and market evaluation capabilities in Go.

## Features

### Core Functionality
- 🚀 **Market Analysis** - Process prematch odds and match results
- 📊 **1X2 Market Evaluation** - Win/Draw/Win predictions
- ⚖️ **Over/Under Goals** - Total goals market analysis
- 🎯 **Correct Score** - Exact score prediction validation
- 💡 **Double Chance** - Bonus market implementation

### Technical Highlights
- 🏗️ Custom JSON unmarshaling for Bet365 data structures
- 📈 Odds conversion system (Decimal ↔ American ↔ Fractional)
- 📉 Value betting identification algorithms
- 📁 Modular architecture for market processors

## Installation

### Prerequisites
- Go 1.18+

```bash
# Clone repository
git clone https://github.com/yesetoda/bet365-evaluator-go.git
cd bet365-evaluator-go

# Install dependencies
go mod tidy

# Build executable
go build main.go

# Run the excutable
./main.go
```

## Usage

```bash
# Basic analysis with default files
go run main.go
```

## Implemented Markets

### 1. Win/Draw/Win (1X2)
```json
"to_win_the_match": {
  "id": "103",
  "name": "1X2",
  "odds": [
    {"id": "1", "name": "1", "odds": "1.80"},
    {"id": "X", "name": "X", "odds": "3.50"},
    {"id": "2", "name": "2", "odds": "4.20"}
  ]
}
```

### 2. Over/Under Goals
```json
"total_goals": {
  "id": "115",
  "name": "Over/Under 2.5 Goals",
  "odds": [
    {"header": "Over", "name": "2.5", "odds": "1.95"},
    {"header": "Under", "name": "2.5", "odds": "1.85"}
  ]
}
```

### 3. Correct Score
```json
"correct_score": {
  "id": "121",
  "name": "Correct Score",
  "odds": [
    {"name": "1-0", "odds": "8.00"},
    {"name": "2-1", "odds": "9.50"},
    {"name": "0-0", "odds": "11.00"}
  ]
}
```

### 4. Double Chance (Bonus)
```json
"double_chance": {
  "id": "131",
  "name": "Double Chance",
  "odds": [
    {"name": "1X", "odds": "1.30"},
    {"name": "X2", "odds": "1.45"},
    {"name": "12", "odds": "1.20"}
  ]
}
```

## Sample Output

```text
==================== MATCH ANALYSIS REPORT ====================
Match: Manchester United vs Chelsea
Date: 2023-08-15 19:45:00
Final Score: 2-1

MARKET EVALUATIONS:
1. 1X2 Market
   Selection: Manchester United @ 1.80 → WIN ✔
   Stake: £100.00 → Return: £180.00 (+£80.00)

2. Over/Under 2.5 Goals
   Selection: Over 2.5 @ 1.95 → WIN ✔
   Stake: £100.00 → Return: £195.00 (+£95.00)

3. Correct Score
   Selection: 2-1 @ 9.50 → WIN ✔
   Stake: £100.00 → Return: £950.00 (+£850.00)

4. Double Chance
   Selection: 1X @ 1.30 → LOSS ✘
   Stake: £100.00 → Return: £0.00 (-£100.00)

SUMMARY:
Total Bets: 4        Winning Bets: 3 (75.00%)
Total Stake: £400.00      Total Return: £1,325.00
Net Profit: £925.00       ROI: 231.25%
===============================================================
```

## Project Structure

```
├── data/                 # Sample JSON files
│   ├── prematch.json     # Prematch odds data
│   └── result.json       # Match result data
├── excuter/              # excuter for cricket and volleyball bets
│   ├── cricket_excuter/cricket_excuter.go  
│   └── volleyball_excuter/volleyball_excuter.go    
├── helpers/              # Core logic
│   ├── cricket_helper/cricket_helper.go 
│   └── volleyball_helper/volleyball_helper.go    
├── models/               # Data structures
│   ├── cricket
│   │   ├── cricket.go
│   │   ├── prematch.go
│   │   └── result.go
│   └── volleyball
│       ├──  prematch.go
│       └──  result.go
└── main.go               # CLI entry point
```

## Key Implementation Details

1. **Custom Unmarshaling**  
   Handles Bet365's inconsistent field naming through custom JSON unmarshalers

2. **Market Processors**  
   Modular design allowing easy addition of new markets:
   ```go
   type MarketProcessor interface {
       Process(prematchData Prematch) []BetSelection
       Evaluate(selection BetSelection, resultData Result) Evaluation
   }
   ```

3. **Odds Conversion**  
   Comprehensive odds formatting system:
   ```go
   // Convert decimal odds to American format
   func DecimalToAmerican(decimal float64) string {
       if decimal >= 2.0 {
           return fmt.Sprintf("+%.0f", (decimal-1)*100)
       }
       return fmt.Sprintf("%.0f", -100/(decimal-1))
   }
   ```
