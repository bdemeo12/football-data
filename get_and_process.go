package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const endpoint = "https://api.football-data.org/v4/competitions?area=2077"

type footballData struct {
	Area areaData
}

type areaData struct {
	Id   string
	Name string
}

func get2024Matches(endpoint string) (*footballData, error) {

	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	godotenv.Load(".env")
	req.Header.Set("X-Auth-Token", os.Getenv("AUTH_TOKEN"))
	req.Header.Set("Content-Type", "application.json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	if err != nil {
		return nil, err
	}

	data := footballData{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func main() {
	data, err := get2024Matches(endpoint)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(data)
	}

}
