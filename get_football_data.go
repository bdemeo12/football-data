package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const endpoint = "https://api.football-data.org/v4/matches"

type footballData struct {
	dateStatus string
	opponenets string
	score      string
}

func getFootballData(endpoint string) (*footballData, error) {

	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application.json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	data := footballData{}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

func main() {
	data, err := getFootballData(endpoint)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(data)
	}

}
