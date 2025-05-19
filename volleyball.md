#Volleyball Betting Data Structures

## Result Data Structures (`result.go`)

### ResultData
- **Success**: Indicates if the API call was successful (1 for success)
- **Results**: Array of MatchResult objects containing actual match data

### MatchResult
- **ID**: Unique identifier for the match
- **SportID**: Identifier for the sport (91 for volleyball)
- **Time**: Unix timestamp of match start time
- **TimeStatus**: Match status (3 indicates completed match)
- **League**: League information (LeagueInfo)
- **Home**: Home team information (TeamInfo)
- **Away**: Away team information (TeamInfo)
- **SS**: Final set score (e.g., "2-3" means home team won 2 sets, away won 3)
- **Scores**: Detailed scores for each set (ScoresInfo)
- **Stats**: Match statistics (StatsInfo)
- **Events**: Array of events that occurred during the match (EventInfo)
- **Extra**: Additional match information (ExtraInfo)
- **InplayCreatedAt**: When inplay data was first created
- **InplayUpdatedAt**: Last update time for inplay data
- **ConfirmedAt**: When the result was confirmed
- **Bet365ID**: Bet365's internal match ID

### LeagueInfo
- **ID**: League identifier
- **Name**: League name (e.g., "Portugal Nacional Women")
- **CC**: Country code (e.g., "pt" for Portugal)

### TeamInfo
- **ID**: Team identifier
- **Name**: Team name
- **ImageID**: Reference ID for team logo/image
- **CC**: Country code

### ScoresInfo
Contains SetScore objects for each set (1-5):
- **Set1**, **Set2**, **Set3**, **Set4**, **Set5**: Each has Home and Away scores

### SetScore
- **Home**: Points scored by home team in the set
- **Away**: Points scored by away team in the set

### StatsInfo
- **PointsWonOnServe**: Array where first element is home team's points won on serve, second is away's
- **LongestStreak**: Array where first element is home team's longest streak, second is away's

### EventInfo
- **ID**: Event identifier
- **Text**: Description of the event (e.g., "Set 1 - Race to 5 points")

### ExtraInfo
- **HomePos**: Home team's position in standings
- **AwayPos**: Away team's position in standings
- **BestOfSets**: Maximum number of sets in match (e.g., "5")
- **Round**: Round number in competition

### BetSelection
Represents a betting selection:
- **Market**: Betting market name
- **MarketID**: Market identifier
- **Selection**: What was selected (e.g., "Winner")
- **SelectionID**: Selection identifier
- **Odds**: Decimal odds
- **Handicap**: Handicap value if applicable
- **StakeAmount**: Amount wagered (for simulation)

### EvaluationResult
Result of evaluating a bet:
- **BetSelection**: The original bet
- **IsWin**: Whether the bet won
- **Explanation**: Reason for win/loss
- **ProfitLoss**: Net profit/loss
- **ReturnAmount**: Total return (stake + profit)
- **ImpliedProbability**: Probability implied by the odds

### MatchStatistics
Comprehensive match stats:
- **TotalMatchPoints**: Sum of all points scored
- **TotalSet1Points** to **TotalSet5Points**: Points in each set
- **HomeSetWins**: Number of sets won by home team
- **AwaySetWins**: Number of sets won by away team
- **MatchWinner**: "home" or "away"
- **Set1Winner** to **Set5Winner**: Winner of each set
- **Set1ExtraPoints** to **Set5ExtraPoints**: Whether set went to extra points
- **CorrectSetScore**: Final set score (e.g., "3-2")
- **TotalSets**: Number of sets played
- **MaximumSets**: Possible sets in match (usually 3 or 5)

## Prematch Data Structures (`prematch.go`)

### PrematchData
- **Success**: API call success indicator
- **Results**: Array of PrematchResult objects

### PrematchResult
- **FI**: Bet365's fixture identifier
- **EventID**: Match identifier (matches ResultData's ID)
- **Main**: Main betting markets (MainData)
- **Others**: Additional betting markets (OtherData array)
- **Schedule**: Scheduled betting markets (ScheduleData)

### MainData
- **UpdatedAt**: Last update timestamp
- **Key**: Internal key for the data
- **Sp**: Sports data (SpData)

### OtherData
- **UpdatedAt**: Last update timestamp
- **Sp**: Sports data (SpData)

### SpData
Contains various betting markets:
- **GameLines**: Main markets like winner, handicap, totals
- **CorrectSetScore**: Correct set score predictions
- **MatchTotalOddEven**: Odd/even total points in match
- **Set1Lines**: Set 1 specific markets
- **Set1ToGoToExtraPoints**: Will set 1 go to extra points
- **Set1TotalOddEven**: Odd/even points in set 1

### MarketData
- **ID**: Market identifier
- **Name**: Market name
- **Odds**: Available odds (OddsData array)
- **Open**: Whether market is open (1/0)

### OddsData
- **ID**: Odds identifier
- **Odds**: Decimal odds value
- **Name**: Selection name
- **Header**: Indicates home (1) or away (2) team
- **Handicap**: Handicap value if applicable

### ScheduleData
- **UpdatedAt**: Last update timestamp
- **Key**: Internal key
- **Sp**: Scheduled odds (ScheduleSp)

### ScheduleSp
- **Main**: Array of scheduled odds (ScheduleOddsData)

### ScheduleOddsData
- **ID**: Odds identifier
- **Odds**: Decimal odds value
- **Name**: Market name
- **Handicap**: Handicap value if applicable

## Observations

1. The data is divided into pre-match (odds) and post-match (results) structures
2. Bet365 uses numeric IDs extensively for teams, matches, markets, and selections
3. The handicap system is represented with positive/negative values
4. Markets include standard options (winner, handicap, totals) and volleyball-specific options (set scores, extra points)
5. Time fields use Unix timestamps
6. The structure allows for multiple markets and sub-markets with different update times
