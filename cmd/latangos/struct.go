package latangos

import "time"

type Flight struct {
	FlightNumber string    `json:"flightNumber"`
	Origin       string    `json:"origin"`
	Destination  string    `json:"destination"`
	Departure    time.Time `json:"departure"`
	Arrival      time.Time `json:"arrival"`
	BasePrice    float64   `json:"basePrice"`
}
