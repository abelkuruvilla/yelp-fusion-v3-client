package yelp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

// AuthOptions contains authentication options for the Yelp API
type ClientOptions struct {
	APIKey string `json:"api_key"`
	Debug  bool   `json:"debug"`
}

// AuthOptionsFromJsonFile returns an ClientOptions object populated
// from a JSON file
func AuthOptionsFromJsonFile(configPath string) (ClientOptions, error) {
	authOptions := ClientOptions{}
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
			Options: ClientOptions{
				APIKey: yelpKey,
				Debug:  dbug,
			},
		}, nil
	}

	// No debug env variable
	return &Client{
		Options: ClientOptions{
			APIKey: yelpKey,
		},
	}, nil
}
