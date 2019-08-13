package yelp

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	yelpURL = "https://api.yelp.com/v3"
)

// Client provides methods to perform requests on the Yelp API
type Client struct {
	AuthOptions AuthOptions
	Debug       bool
}

// NewClientCustom creates a new Client instance
func NewClientCustom(authOptions AuthOptions) Client {
	return Client{
		AuthOptions: authOptions,
	}
}

// NewClient creates a new client instance from the environment variables or a config.json file
// set the variables YELP_API_KEY and YELP_DEBUG
// However if a config.json file is present with the key 'api_key', it will over-ride the environment key values
func NewClient() (*Client, error) {

	client, errEnv := clientFromEnvVariables()

	ao, errFile := AuthOptionsFromJsonFile("./config.json")
	if errFile == nil {

		if errEnv == nil {
			// Return client from file variables
			client.AuthOptions = ao
			return client, nil
		}
		// Error fetching from env
		return &Client{AuthOptions: ao}, nil

	}
	if errEnv == nil {
		return client, nil

	}
	return nil, fmt.Errorf("Error fetching configuration from environment and file \n Env error :%v \n File Error :%v", errEnv, errFile)

}

func (client *Client) request(method string, endpoint string, params map[string]interface{}, response interface{}) error {
	url := fmt.Sprintf("%s%s", yelpURL, endpoint)
	if client.Debug {
		log.Printf("%s %s %+v\n", method, url, params)
	}
	httpClient := &http.Client{}
	paramsAsBytes, err := json.Marshal(params)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(paramsAsBytes))
	if err != nil {
		return err
	}
	req.ContentLength = int64(len(paramsAsBytes))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", client.AuthOptions.APIKey))
	res, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}

		// If status moved error
		if res.StatusCode == http.StatusMovedPermanently {

			var eb errorBody
			if err := json.Unmarshal(data, &eb); err != nil {
				return err
			}
			return eb.Error
		}

		return errors.New(string(data))
	}
	json.NewDecoder(res.Body).Decode(response)
	return nil
}

func (client *Client) query(url string, request interface{}, response interface{}) error {
	method := "GET"

	query, err := structToQueryParams(request)
	if err != nil {
		return err
	}
	endpoint := url + query

	params := make(map[string]interface{})

	if err := client.request(method, endpoint, params, response); err != nil {
		return err
	}
	return nil
}
