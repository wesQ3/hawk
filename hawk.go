package hawk

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type RegisterRequest struct {
	Faction string `json:"faction"`
	Symbol  string `json:"symbol"`
}

func Register(faction, symbol string) (*http.Response, error) {
	url := "https://api.spacetraders.io/v2/register"

	requestData := &RegisterRequest{
		Faction: faction,
		Symbol:  symbol,
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type ShipsResponse struct {
	Data []Ship `json:"data"`
	Meta Meta   `json:"meta"`
}

type Meta struct {
	Total int `json:"total"`
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

func Ships(token string) ([]Ship, error) {
	url := "https://api.spacetraders.io/v2/my/ships"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var shipsResponse ShipsResponse
	err = json.Unmarshal(body, &shipsResponse)
	if err != nil {
		return nil, err
	}

	return shipsResponse.Data, nil
}
