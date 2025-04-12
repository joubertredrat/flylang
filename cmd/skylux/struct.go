package skylux

import "time"

type (
	FlightRoute struct {
		From string `json:"from"`
		To   string `json:"to"`
	}

	Flight struct {
		Callsign  string        `json:"callsign"`
		Routes    []FlightRoute `json:"routes"`
		Departure time.Time     `json:"departure"`
		Arrival   time.Time     `json:"arrival"`
		SeatPrice float64       `json:"seatPrice"`
	}
)
