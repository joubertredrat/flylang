package aeronyx

import "time"

type FlightLeg struct {
	FlightNumber    string         `json:"flight_number"`
	From            string         `json:"from"`
	To              string         `json:"to"`
	BasePrice       float64        `json:"base_price"`
	OperatingDays   []time.Weekday `json:"operating_days"`
	DepartureHour   int            `json:"departure_hour"`
	DepartureTime   string         `json:"departure_time"`
	DurationMinutes int            `json:"duration_minutes"`
}

type Itinerary struct {
	Legs       []FlightLeg `json:"legs"`
	TotalPrice float64     `json:"total_price"`
}

type FlightResponse struct {
	FlightNumber         string `json:"flight_number"`
	From                 string `json:"from"`
	To                   string `json:"to"`
	TotalPrice           string `json:"total_price"`
	DepartureTime        string `json:"departure_time"`
	TotalDurationMinutes int    `json:"total_duration_minutes"`
	NumberOfConnections  int    `json:"number_of_connections"`
}
