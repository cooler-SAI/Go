package main

import (
	"encoding/json"
	"fmt"
)

type Bird struct {
	Species string `json:"Knurl"`
}

type ManInfo struct {
	Name     string   `json:"name"`
	Age      int      `json:"age"`
	Email    string   `json:"email"`
	IsActive bool     `json:"isActive"`
	Roles    []string `json:"roles"`
	Address  struct {
		Street     string `json:"street"`
		City       string `json:"city"`
		State      string `json:"state"`
		PostalCode string `json:"postalCode"`
	} `json:"address"`
}

type Country struct {
	Country         string `json:"country"`
	DisplayName     string `json:"display_name"`
	Email           string `json:"email"`
	ExplicitContent struct {
		FilterEnabled bool `json:"filter_enabled"`
		FilterLocked  bool `json:"filter_locked"`
	} `json:"explicit_content"`
	ExternalUrls struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
	Followers struct {
		Href  string `json:"href"`
		Total int    `json:"total"`
	} `json:"followers"`
	Href   string `json:"href"`
	ID     string `json:"id"`
	Images []struct {
		URL    string `json:"url"`
		Height int    `json:"height"`
		Width  int    `json:"width"`
	} `json:"images"`
	Product string `json:"product"`
	Type    string `json:"type"`
	URI     string `json:"uri"`
}

func main() {
	// Marshal Bird struct to JSON
	pigeon := &Bird{
		Species: "Pigeon",
	}

	data, err := json.Marshal(pigeon)
	if err != nil {
		fmt.Println("Error marshalling Bird struct:", err)
		return
	}

	fmt.Println("Marshalled Bird JSON:", string(data))

	jsonStr := `{
		"name": "John Doe",
		"age": 30,
		"email": "john.doe@example.com",
		"isActive": true,
		"roles": ["user", "admin"],
		"address": {
			"street": "123 Main St",
			"city": "London",
			"state": "CA",
			"postalCode": "12345"
		}
	}`

	var man ManInfo
	err = json.Unmarshal([]byte(jsonStr), &man)
	if err != nil {
		fmt.Println("Error unmarshalling JSON string:", err)
		return
	}

	fmt.Printf("Unmarshalled ManInfo struct: %+v\n", man)

	country := &Country{
		Country:     "UK",
		DisplayName: "United Kingdom",
		Email:       "info@uk.gov",
		ExplicitContent: struct {
			FilterEnabled bool `json:"filter_enabled"`
			FilterLocked  bool `json:"filter_locked"`
		}{
			FilterEnabled: true,
			FilterLocked:  false,
		},
		ExternalUrls: struct {
			Spotify string `json:"spotify"`
		}{
			Spotify: "https://spotify.com/uk",
		},
		Followers: struct {
			Href  string `json:"href"`
			Total int    `json:"total"`
		}{
			Href:  "",
			Total: 1000000,
		},
		Href: "https://example.com/uk",
		ID:   "uk",
		Images: []struct {
			URL    string `json:"url"`
			Height int    `json:"height"`
			Width  int    `json:"width"`
		}{
			{
				URL:    "https://example.com/image1.jpg",
				Height: 100,
				Width:  100,
			},
		},
		Product: "premium",
		Type:    "country",
		URI:     "spotify:country:uk",
	}

	countryData, err := json.Marshal(country)
	if err != nil {
		fmt.Println("Error marshalling Country struct:", err)
		return
	}

	fmt.Println("Marshalled Country JSON:", string(countryData))
}
