package yelp

import (
	"encoding/json"
	"fmt"
)

// Business represents a business object from the Yelp API
type Business struct {
	ID           string         `json:"id,omitempty"`
	Alias        string         `json:"alias,omitempty"`
	Name         string         `json:"name,omitempty"`
	ImageUrl     string         `json:"image_url,omitempty"`
	IsClaimed    bool           `json:"is_claimed,omitempty"`
	IsClosed     bool           `json:"is_closed,omitempty"`
	Url          string         `json:"url,omitempty"`
	Phone        string         `json:"phone,omitempty"`
	DisplayPhone string         `json:"display_phone,omitempty"`
	ReviewCount  uint           `json:"review_count,omitempty"`
	Categories   []Category     `json:"categories,omitempty"`
	Rating       float64        `json:"rating,omitempty"`
	Location     Location       `json:"location,omitempty"`
	Coordinates  Coordinate     `json:"coordinates,omitempty"`
	Photos       []string       `json:"photos,omitempty"`
	Price        string         `json:"price,omitempty"`
	Hours        []Hour         `json:"hours,omitempty"`
	Transactions []string       `json:"transactions,omitempty"`
	Messaging    Messaging      `json:"messaging,omitempty"`
	SpecialHours []SpecialHour  `json:"special_hours,omitempty"`
	Distance     float64        `json:"distance,omitempty"`
	Reviews      ReviewResponse `json:"reviews,omitempty"`
}

type BusinessDetailParams struct {
	Locale string `json:"locale,omitempty"`
}

type Messaging struct {
	URL         string `json:"url,omitempty"`
	UsecaseText string `json:"use_case_text,omitempty"`
}

type BusinessMatchParams struct {
	Locale   string `json:"locale,omitempty"`
	Name     string `json:"name,omitempty"`
	Address1 string `json:"address1,omitempty"`
}

type BusinessSearchParams struct {
	Locale     string  `json:"locale,omitempty"`
	Term       string  `json:"term,omitempty"`
	Location   string  `json:"location,omitempty"`
	Latitude   float64 `json:"latitude,omitempty"`
	Longitude  float64 `json:"longitude,omitempty"`
	Radius     uint    `json:"radius,omitempty"`
	Categories string  `json:"categories,omitempty"`
	Limit      uint    `json:"limit,omitempty"`
	Offset     uint    `json:"offset,omitempty"`
	SortBy     string  `json:"sort_by,omitempty"`
	Price      string  `json:"price,omitempty"`
	OpenNow    bool    `json:"open_now,omitempty"`
	OpenAt     int     `json:"open_at,omitempty"`
	Attributes string  `json:"attributes,omitempty"`
}

type BusinessSearchResponse struct {
	Total      uint       `json:"total,omitempty"`
	Businesses []Business `json:"businesses,omitempty"`
}

type PhoneSearchParams struct {
	Phone  string `json:"phone,omitempty"`
	Locale string `json:"locale,omitempty"`
}

type PhoneSearchResponse struct {
	Total      uint       `json:"total,omitempty"`
	Businesses []Business `json:"businesses,omitempty"`
}

type TransactionSearchParam struct {
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
	Location  string  `json:"location,omitempty"`
}

type ErrBusinessMigrated struct {
	Code          string `json:"code,omitempty"`
	Description   string `json:"description,omitempty"`
	NewBusinessID string `json:"new_business_id,omitempty"`
}

func (err ErrBusinessMigrated) Error() string {
	b, _ := json.Marshal(err)
	return string(b)
}

type errorBody struct {
	Error ErrBusinessMigrated `json:"error,omitempty"`
}

// Search sends a business search request to the Yelp API
// See https://www.yelp.com/developers/documentation/v3/business_search
func (client *Client) Search(request BusinessSearchParams) (response BusinessSearchResponse, err error) {

	url := fmt.Sprintf("/businesses/search")
	if err := client.query(url, &request, &response); err != nil {
		return response, err
	}
	return response, nil

}

// SearchPhone is used to search businesses with a particular phone number.
// See https://www.yelp.com/developers/documentation/v3/business_search_phone
func (client *Client) SearchPhone(request PhoneSearchParams) (response PhoneSearchResponse, err error) {

	url := fmt.Sprintf("/businesses/search/phone")
	if err := client.query(url, &request, &response); err != nil {
		return response, err
	}
	return response, nil

}

// GetBusinessDetails is used to get information regarding a business from yelp, given its id.
// We also fetch the reviews of that business from the reviews API.
// See https://www.yelp.com/developers/documentation/v3/business - Business Details API
// See https://www.yelp.com/developers/documentation/v3/business_reviews - Reviews API
func (client *Client) GetBusinessDetails(id string, request BusinessDetailParams) (response Business, err error) {

	url := fmt.Sprintf("/businesses/%s", id)
	err = client.query(url, &request, &response)
	if err != nil {
		return response, err
	}

	// Attach Reviews too
	err = response.addReviews(client, ReviewParams{
		Locale: request.Locale,
	})
	if err != nil {
		return response, err
	}
	return response, nil
}

// BusinessTransactionSearch is used to get businesses that support the specified transaction.
// See https://www.yelp.com/developers/documentation/v3/transaction_search
func (client *Client) BusinessTransactionSearch(transactionType string, request TransactionSearchParam) (response BusinessSearchResponse, err error) {

	url := fmt.Sprintf("/transactions/%s/search", transactionType)
	err = client.query(url, &request, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}
