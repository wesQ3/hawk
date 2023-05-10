package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/wesQ3/hawk"
)

type RegisterResponse struct {
	Token string `json:"token"`
}


func main() {
	faction := "COSMIC"
	symbol := "HAWK2341"

	resp, err := hawk.Register(faction, symbol)
	if err != nil {
		log.Fatalf("Error registering: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusCreated {
		var registerResponse RegisterResponse
		err = json.NewDecoder(resp.Body).Decode(&registerResponse)
		if err != nil {
			log.Fatalf("Error decoding JSON response: %v", err)
		}

		fmt.Println("Registration successful")
		fmt.Printf("HAWK_TOKEN=%s\n", registerResponse.Token)
	} else {
		var errorResponse map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&errorResponse)
		if err != nil {
			log.Fatalf("Error decoding JSON error response: %v", err)
		}

		errorResponseJSON, err := json.MarshalIndent(errorResponse, "", "  ")
		if err != nil {
			log.Fatalf("Error formatting JSON error response: %v", err)
		}

		fmt.Printf("Registration failed with status code: %d\n", resp.StatusCode)
		fmt.Println("Error response:")
		fmt.Println(string(errorResponseJSON))
	}
}
