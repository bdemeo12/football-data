package main

import (
	"fmt"
	"os"

	getandprocess "github.com/bdemeo12/football-data/get-and-process"
	"github.com/joho/godotenv"
)

func main() {

	// get envars
	godotenv.Load("get-and-process/.env")
	endpoint := os.Getenv("ENDPOINT")
	authToken := os.Getenv("AUTH_TOKEN")

	data, err := getandprocess.Get2024Matches(endpoint, authToken)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(data.Matches[0])
}
