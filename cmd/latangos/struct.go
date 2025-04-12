package latangos

import "time"

type Flight struct {
	FlightNumber string    `json:"flight_number"`
	Origin       string    `json:"origin"`
	Destination  string    `json:"destination"`
	Departure    time.Time `json:"departure"`
	Arrival      time.Time `json:"arrival"`
	BasePrice    float64   `json:"base_price"`
}
