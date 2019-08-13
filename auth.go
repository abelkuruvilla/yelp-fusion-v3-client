package yelp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

// AuthOptions contains authentication options for the Yelp API
type AuthOptions struct {
	APIKey string `json:"api_key"`
}

// AuthOptionsFromJsonFile returns an AuthOptions object populated
// from a JSON file
func AuthOptionsFromJsonFile(configPath string) (AuthOptions, error) {
	authOptions := AuthOptions{}
	bytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		return authOptions, err
	}
	err = json.Unmarshal(bytes, &authOptions)
	if err != nil {
		return authOptions, err
	}
	return authOptions, nil
}

func clientFromEnvVariables() (*Client, error) {
	yelpKey, isPresent := os.LookupEnv("YELP_API_KEY")
	if !isPresent {
		return nil, fmt.Errorf("YELP_API_KEY is not set in the environment variable")
	}

	debug, isPresent := os.LookupEnv("YELP_DEBUG")
	if isPresent {
		dbug, err := strconv.ParseBool(debug)
		if err != nil {
			return nil, err
		}
		return &Client{
			AuthOptions: AuthOptions{
				APIKey: yelpKey,
			},
			Debug: dbug,
		}, nil
	}

	// No debug env variable
	return &Client{
		AuthOptions: AuthOptions{
			APIKey: yelpKey,
		},
	}, nil
}
