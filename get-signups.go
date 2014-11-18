package main

// This script prints some (1 page) of the people in a nation.

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	NationName  = "NATION_NAME"
	AccessToken = "ACCESS_TOKEN"
)

// Returns a URL for the people index endpoint
func peopleUrl() string {
	return fmt.Sprintf("https://%s.nationbuilder.com/api/v1/people", NationName)
}

// Returns the authorization header, which includes the access token
func authorizationHeader() string {
	return fmt.Sprintf("Bearer %s", AccessToken)
}

// Contains the results of a request to the people index endpoint
type Results struct {
	People []struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	} `json:"results"`
}

func main() {
	// Construct the request
	req, err := http.NewRequest("GET", peopleUrl(), nil)
	req.Header.Add("Authorization", authorizationHeader())
	req.Header.Add("Accept", "application/json")

	// Perform the request
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	// Deserialize the results
	var results Results
	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&results); err != nil {
		panic(err)
	}

	// Print out the results
	fmt.Printf("People in nation %s:\n", NationName)
	for _, person := range results.People {
		fmt.Printf("    %s %s\n", person.FirstName, person.LastName)
	}
}
