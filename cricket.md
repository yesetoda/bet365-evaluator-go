# Cricket Betting Data Structures

## Core Data Structures (`cricket.go`)

### **Odd**
- **ID**: Unique identifier for the odds entry (e.g., `PC123456`)
- **Odds**: Decimal odds value as a string (e.g., `"1.83"`)
- **Name**: Selection name (e.g., `"Over 6.5 runs"`)
- **Name2**: Alternative name (optional)
- **Header**: Indicates team context (`1` = home, `2` = away)
- **Handicap**: Handicap value (e.g., `"+2.5"` for handicap markets)

### **BetSelection**
- **Market**: Betting category (e.g., `"Match Winner"`)
- **Selection**: Chosen option (e.g., `"Mumbai Indians"`)
- **Odds**: Decimal odds at the time of bet placement
- **IsWinner**: Outcome status (`true`/`false`)
- **Evaluation**: Explanation of bet result
- **ConfidenceLevel**: Risk level (e.g., `"High"`)
- **AvailableOptions**: All market choices (e.g., `["Over 6.5", "Under 6.5"]`)
- **MarketDescription**: Rules/context for the market
- **PotentialProfit**: Calculated profit if the bet wins
- **RiskAssessment**: Risk category (`Low`/`Medium`/`High`)
- **OddsDecimal/American/Fractional**: Odds in different formats

### **DetailedMatchInfo**
- **HomeTeam/AwayTeam**: Team names (e.g., `"Rajasthan Royals"`)
- **HomeScore/AwayScore**: Numerical scores (e.g., `217` vs. `117`)
- **Stadium**: Venue name (e.g., `"Sawai Mansingh Stadium"`)
- **City/Country**: Location details (e.g., `"Jaipur, India"`)
- **Capacity**: Stadium capacity (e.g., `"23,185"`)
- **MatchDate**: Date in `YYYY-MM-DD` format
- **LeagueName**: Tournament name (e.g., `"Indian Premier League"`)
- **BattingStats**: Map of player batting metrics (see `BattingStats`)
- **BowlingStats**: Map of player bowling metrics (see `BowlingStats`)

### **BattingStats**
- **Runs**: Total runs scored
- **Balls**: Total balls faced
- **StrikeRate**: Runs per 100 balls (e.g., `150.0`)
- **Boundaries**: Number of 4s hit
- **Sixes**: Number of 6s hit

### **BowlingStats**
- **Overs**: Overs bowled (e.g., `4.2` = 4 overs + 2 balls)
- **RunsConceded**: Runs given away
- **Wickets**: Wickets taken
- **Economy**: Runs conceded per over (e.g., `6.5`)

### **BettingHistory**
- **Market**: Market type (e.g., `"Player Performance"`)
- **WinPercentage**: Historical success rate (e.g., `65.2%`)
- **AvgOdds**: Average odds used
- **TotalBets**: Number of bets placed
- **ProfitLoss**: Net profit/loss (e.g., `+$320.50`)

---

## Prematch Data Structures (`prematch.go`)

### **CricketPrematchData**
- **Success**: API status (`1` = success, `0` = failure)
- **Results**: Array of prematch markets for a match

#### Nested Structures:
**1. First Over Markets**  
- `1st_over_total_runs`: Total runs in the first over (e.g., `Over/Under 6.5`)
- `1st_over_total_runs_odd_even`: Odd/Even runs in the first over

**2. Innings Markets**  
- `1st_innings_score`: Predicted score ranges for the first innings
- `1st_innings_of_match_bowled_out?`: Whether the team will be bowled out

**3. Main Markets**  
- `to_win_the_match`: Match winner odds
- `team_top_batter/bowler`: Odds for top performers
- `player_of_the_match`: Player of the match predictions
- `1st_wicket_method`: How the first wicket will fall (e.g., `"Caught"`)

**4. Match Specials**  
- `team_to_make_highest_1st_6_overs_score`: Powerplay leader
- `to_go_to_super_over?`: Odds for a tied match requiring a super over
- `a_fifty/hundred_to_be_scored`: Milestone markets

**5. Player Markets**  
- `batter_match_runs`: Projected runs for specific batters
- `batter_milestones`: Odds for a player to score 50/100
- `bowler_total_match_wickets`: Projected wickets for bowlers

---

## Result Data Structures (`result.go`)

### **CricketResultData**
- **Success**: API status (`1` = success)
- **Results**: Array of completed matches

#### Match Details:
- **ID**: Unique match identifier (e.g., `"9703206"`)
- **SportID**: Sport type (`3` = cricket)
- **TimeStatus**: Match status (`3` = completed)
- **League**:  
  - **ID**: League ID (e.g., `"26431"` for IPL)  
  - **Name**: League name  
  - **CC**: Country code (e.g., `"in"` for India)
- **Home/Away**:  
  - **ID**: Team ID  
  - **Name**: Team name  
  - **ImageID**: ID for team logo  
  - **CC**: Team country code
- **SS**: Score format `"117-217"` (Home-Away runs)
- **Extra**:  
  - `stadium_data`:  
    - **ID**: Venue ID  
    - **Name**: Stadium name  
    - **City/Country**: Location  
    - **Capacity**: Seating capacity  
    - **GoogleCoo**: GPS coordinates for maps
- **HasLineup**: `1` if lineup data exists
- **Inplay_*_at**: Unix timestamps for match lifecycle events
- **Bet365ID**: Bet365’s internal match ID

---

## Observations

1. **ID System**: All entities (matches, teams, markets) use numeric IDs (e.g., `"9703206"`).
2. **Score Format**:  
   - `ss` field uses `"HomeScore-AwayScore"` (e.g., `"117-217"`).
3. **Market Structure**:  
   - Markets are nested under `sp` (sports properties).  
   - Example: `1st_wicket_method` under `main.sp`.
4. **Status Indicators**:  
   - `time_status`: `1` = upcoming, `2` = live, `3` = completed.  
   - `open`: `1` = market active, `0` = closed.
5. **Geo Data**: Stadiums include Google Maps coordinates (`googlecoords`).
6. **Player Focus**: Specialized markets for batters/bowlers (e.g., `batter_milestones`).
7. **Unix Timestamps**: All time fields (e.g., `inplay_created_at`) use Unix epoch format.