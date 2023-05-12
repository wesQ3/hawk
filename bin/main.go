package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/wesQ3/hawk"
)

func main() {
	faction := "COSMIC"
	symbol := "HAWK0304"
	token := os.Getenv("HAWK_TOKEN")

	if token == "" {
		registerSymbol(faction, symbol)
	}

	resp, err := hawk.Ships(token)
	if err != nil {
		log.Fatalf("Error ships: %v", err)
	}
	fmt.Printf("ships: %s\n", resp)
}

type RegisterResponse struct {
	Token string `json:"token"`
}

func registerSymbol(faction, symbol string) {
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
