package hawk

import (
	"bytes"
	"encoding/json"
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
