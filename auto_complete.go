package yelp

import "fmt"

type AutoCompleteParams struct {
	Locale    string  `json:"locale,omitempty"`
	Text      string  `json:"text,omitempty"`
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
}

type AutoCompleteResponse struct {
	Terms      []Term     `json:"terms,omitempty"`
	Businesses []Business `json:"businesses,omitempty"`
	Categories []Category `json:"categories,omitempty"`
}

type Term struct {
	Text string `json:"text,omitempty"`
}

// AutoComplete is used to return auto-complete suggestions for a term from yelp.
// See https://www.yelp.com/developers/documentation/v3/autocomplete
func (client *Client) AutoComplete(request AutoCompleteParams) (response AutoCompleteResponse, err error) {
	url := fmt.Sprintf("/autocomplete")

	err = client.query(url, &request, &response)
	if err != nil {
		return response, err
	}
	return response, nil

}
