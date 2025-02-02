package getandprocess

import "fmt"

type stageScoreData struct {
	HomeTeamWins int
	AwayTeamWins int
	Draws        int
}

func FinalScore(footballData footballData) {
	// map of stage to final score data
	finalScoreData := make(map[string]stageScoreData)

	for _, match := range footballData.Matches {
		value := finalScoreData[match.Stage]

		switch winner := match.Score.Winner; winner {
		case "HOME_TEAM":
			value.HomeTeamWins++
		case "AWAY_TEAM":
			value.AwayTeamWins++
		case "DRAW":
			value.Draws++
		}

		finalScoreData[match.Stage] = value
	}

	fmt.Println("Final Scores Per Stage:")
	for key, value := range finalScoreData {
		fmt.Printf("%s: Home Team Wins: %d Away Team Wins: %d Draws: %d\n", key, value.HomeTeamWins, value.AwayTeamWins, value.Draws)
	}

}
