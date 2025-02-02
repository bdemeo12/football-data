package getandprocess

import (
	"fmt"
	"sort"
)

func TopThreeTeams(footballData footballData) {
	teamData := organizeTeamData(footballData)

	var keys []string
	for key := range teamData {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool { return teamData[keys[i]] > teamData[keys[j]] })

	// we only want the top three
	keys = keys[:3]

	fmt.Println("Top Three Teams:")
	for _, key := range keys {
		fmt.Printf("%s: %d\n", key, teamData[key])
	}

}

func organizeTeamData(footballData footballData) map[string]int {
	// Map of team name to goals scored
	teamData := make(map[string]int)

	for _, match := range footballData.Matches {

		// Home Team
		homeValue := teamData[match.HomeTeam.Name]
		homeValue += match.Score.Fulltime.Home
		teamData[match.HomeTeam.Name] = homeValue

		// Away Team
		awayValue := teamData[match.AwayTeam.Name]
		awayValue += match.Score.Fulltime.Away
		teamData[match.AwayTeam.Name] = awayValue
	}

	return teamData
}
