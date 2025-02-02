package getandprocess

import (
	"encoding/json"
	"io"
	"net/http"
)

func Get2024Matches(endpoint string, authToken string) (footballData, error) {

	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return footballData{}, err
	}

	req.Header.Set("X-Auth-Token", authToken)
	req.Header.Set("Content-Type", "application.json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return footballData{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return footballData{}, err
	}

	data := footballData{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return footballData{}, err
	}

	return data, nil
}
