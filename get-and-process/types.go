package getandprocess

type footballData struct {
	Matches []matchData
}

type matchData struct {
	HomeTeam homeTeam
	AwayTeam awayTeam
	Score    score
	Stage    string
}

type homeTeam struct {
	Name string
}

type awayTeam struct {
	Name string
}

type score struct {
	Winner   string
	Fulltime fullTime
}

type fullTime struct {
	Home int
	Away int
}
