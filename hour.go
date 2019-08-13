package yelp

import "time"

type Hour struct {
	IsOpenNow bool     `json:"is_open_now,omitempty"`
	HoursType string   `json:"hours_type,omitempty"`
	Open      OpenTime `json:"open,omitempty"`
}
type OpenTime struct {
	Day         uint   `json:"day,omitempty"`
	Start       string `json:"start,omitempty"`
	End         string `json:"end,omitempty"`
	IsOverNight bool   `json:"is_overnight,omitempty"`
}

type SpecialHour struct {
	Date        time.Time `json:"time,omitempty"`
	IsClosed    bool      `json:"is_closed,omitempty"`
	Start       string    `json:"start,omitempty"`
	End         string    `json:"end,omitempty"`
	IsOverNight bool      `json:"is_overnight,omitempty"`
}
