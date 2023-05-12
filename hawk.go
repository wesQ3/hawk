package hawk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type RegisterRequest struct {
	Faction string `json:"faction"`
	Symbol  string `json:"symbol"`
}

type RegisterResponse struct {
	Data struct {
		Token string `json:"token"`
	} `json:"data"`
}

func Register(faction, symbol string) (*RegisterResponse, error) {
	url := "https://api.spacetraders.io/v2/register"

	payload := map[string]string{
		"faction": faction,
		"symbol":  symbol,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
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

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("non-201 response: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var registerResponse RegisterResponse
	err = json.Unmarshal(body, &registerResponse)
	if err != nil {
		return nil, err
	}

	return &registerResponse, nil
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
