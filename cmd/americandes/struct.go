package americandes

import "time"

type (
	Flight struct {
		ID                string    `json:"id"`
		IATACodeFrom      string    `json:"iata_code_from"`
		IATACodeTo        string    `json:"iata_code_to"`
		DepartureTime     time.Time `json:"departure_time"`
		DurationInMinutes int       `json:"duration_in_minutes"`
		Amount            int       `json:"amount"`
	}

	FlightsResponse struct {
		Metadata FlightsResponseMetadata `json:"metadata"`
		Flights  []Flight                `json:"flights"`
	}

	FlightsResponseMetadata struct {
		TotalFlights int                           `json:"total_flights"`
		ServerTime   string                        `json:"server_time"`
		Search       FlightsResponseMetadataSearch `json:"search"`
	}

	FlightsResponseMetadataSearch struct {
		DepartureDay   string `json:"departure_day"`
		DepartureMonth string `json:"departure_month"`
		DepartureYear  string `json:"departure_year"`
		CodeFrom       string `json:"code_from"`
		CodeTo         string `json:"code_to"`
	}
)
