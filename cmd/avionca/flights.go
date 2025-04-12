package avionca

import (
	"time"
)

func flights(origin, destination string, date time.Time) []*Flight {
	switch date.Weekday() {
	case time.Monday:
		return monday(origin, destination, date)
	case time.Tuesday:
		return tuesday(origin, destination, date)
	case time.Wednesday:
		return wednesday(origin, destination, date)
	case time.Thursday:
		return thursday(origin, destination, date)
	case time.Friday:
		return friday(origin, destination, date)
	case time.Saturday:
		return saturday(origin, destination, date)
	case time.Sunday:
		return sunday(origin, destination, date)
	default:
		return []*Flight{}
	}
}

func monday(origin, destination string, date time.Time) []*Flight {
	return []*Flight{
		{
			FlightNumber: "AV001",
			Origin:       origin,
			Destination:  destination,
			Departure:    date.Add(2 * time.Hour).Format(time.RFC3339),
			Duration:     180, // 3 hours in minutes
			Price:        250.00,
		},
	}
}

func tuesday(origin, destination string, date time.Time) []*Flight {
	return []*Flight{
		{
			FlightNumber: "AV002",
			Origin:       origin,
			Destination:  destination,
			Departure:    date.Add(3 * time.Hour).Format(time.RFC3339),
			Duration:     240, // 4 hours in minutes
			Price:        300.00,
		},
	}
}

func wednesday(origin, destination string, date time.Time) []*Flight {
	return []*Flight{
		{
			FlightNumber: "AV003",
			Origin:       origin,
			Destination:  destination,
			Departure:    date.Add(4 * time.Hour).Format(time.RFC3339),
			Duration:     200, // 3 hours 20 minutes in minutes
			Price:        280.00,
		},
	}
}

func thursday(origin, destination string, date time.Time) []*Flight {
	return []*Flight{
		{
			FlightNumber: "AV004",
			Origin:       origin,
			Destination:  destination,
			Departure:    date.Add(5 * time.Hour).Format(time.RFC3339),
			Duration:     300, // 5 hours in minutes
			Price:        350.00,
		},
	}
}

func friday(origin, destination string, date time.Time) []*Flight {
	return []*Flight{
		{
			FlightNumber: "AV005",
			Origin:       origin,
			Destination:  destination,
			Departure:    date.Add(6 * time.Hour).Format(time.RFC3339),
			Duration:     360, // 6 hours in minutes
			Price:        400.00,
		},
	}
}

func saturday(origin, destination string, date time.Time) []*Flight {
	return []*Flight{
		{
			FlightNumber: "AV006",
			Origin:       origin,
			Destination:  destination,
			Departure:    date.Add(7 * time.Hour).Format(time.RFC3339),
			Duration:     420, // 7 hours in minutes
			Price:        450.00,
		},
	}
}

func sunday(origin, destination string, date time.Time) []*Flight {
	return []*Flight{
		{
			FlightNumber: "AV007",
			Origin:       origin,
			Destination:  destination,
			Departure:    date.Add(8 * time.Hour).Format(time.RFC3339),
			Duration:     480, // 8 hours in minutes
			Price:        500.00,
		},
	}
}
