package yelp

import (
	"fmt"
	"time"
)

type Event struct {
	AttendingCount  uint      `json:"attending_count,omitempty"`
	Category        string    `json:"category,omitempty"`
	Cost            float64   `json:"cost,omitempty"`
	CostMax         float64   `json:"cost_max,omitempty"`
	Description     string    `json:"description,omitempty"`
	EventSiteURL    string    `json:"event_site_url,omitempty"`
	ID              string    `json:"id,omitempty"`
	ImageURL        string    `json:"image_url,omitempty"`
	InterestedCount uint      `json:"interested_count,omitempty"`
	IsCanceled      bool      `json:"is_canceled,omitempty"`
	IsFree          bool      `json:"is_free,omitempty"`
	IsOfficial      bool      `json:"is_official,omitempty"`
	Latitude        float64   `json:"latitude,omitempty"`
	Longitude       float64   `json:"longitude,omitempty"`
	Name            string    `json:"name,omitempty"`
	TicketsURL      string    `json:"tickets_url,omitempty"`
	TimeEnd         time.Time `json:"time_end,omitempty"`
	TimeStart       time.Time `json:"time_start,omitempty"`
	Location        Location  `json:"location,omitempty"`
	BusinessID      string    `json:"business_id,omitempty"`
}

type EventGetParam struct {
	Locale string `json:"locale,omitempty"`
}

type EventSearchParam struct {
	Locale         string   `json:"locale,omitempty"`
	Offset         int      `json:"offset,omitempty"`
	Limit          uint     `json:"limit,omitempty"`
	SortBy         string   `json:"sort_by,omitempty"`
	SortOn         string   `json:"sort_on,omitempty"`
	StartDate      int      `json:"start_date,omitempty"`
	EndDate        int      `json:"end_date,omitempty"`
	Categories     string   `json:"categories,omitempty"`
	IsFree         bool     `json:"is_free,omitempty"`
	Location       string   `json:"location,omitempty"`
	Latitude       float64  `json:"latitude,omitempty"`
	Longitude      float64  `json:"longitude,omitempty"`
	Radius         uint     `json:"radius,omitempty"`
	ExcludedEvents []string `json:"excluded_events,omitempty"`
}

type EventSearchResponse struct {
	Total  uint    `json:"total,omitempty"`
	Events []Event `json:"events,omitempty"`
}

type FeaturedEventParam struct {
	Locale    string  `json:"locale,omitempty"`
	Location  string  `json:"location,omitempty"`
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
}

// GetEventDetails is used to fetch details of an event from Yelp API.
// https://www.yelp.com/developers/documentation/v3/event
func (client *Client) GetEventDetails(id string, request EventGetParam) (response Event, err error) {
	url := fmt.Sprintf("/events/%s", id)
	err = client.query(url, &request, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}

// SearchEvent is used to search for events with the Yelp API
// See https://www.yelp.com/developers/documentation/v3/event_search
func (client *Client) SearchEvent(request EventSearchParam) (response EventSearchResponse, err error) {
	url := fmt.Sprintf("/events")
	err = client.query(url, &request, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}

// GetFeaturedEvent is used to search for featured events with the Yelp API
// See https://www.yelp.com/developers/documentation/v3/featured_event
func (client *Client) GetFeaturedEvent(request FeaturedEventParam) (response Event, err error) {
	url := fmt.Sprintf("/events/featured")
	err = client.query(url, &request, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}
