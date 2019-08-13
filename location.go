package yelp

// Location represents a location object from the Yelp API
type Location struct {
	City           string   `json:"city,omitempty"`
	State          string   `json:"state,omitempty"`
	ZipCode        string   `json:"zip_code,omitempty"`
	Country        string   `json:"country,omitempty"`
	Address1       string   `json:"address1,omitempty"`
	Address2       string   `json:"address2,omitempty"`
	Address3       string   `json:"address3,omitempty"`
	CrossStreets   string   `json:"cross_streets,omitempty"`
	DisplayAddress []string `json:"display_address,omitempty"`
}

// Coordinate represents a coordinate object from the Yelp API
type Coordinate struct {
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
}
