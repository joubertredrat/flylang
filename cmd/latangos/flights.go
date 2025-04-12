package latangos

import "time"

func flights(origin, destination string, date time.Time) []Flight {
	switch date.Weekday() {
	case time.Monday:
		return monday(origin, destination)
	case time.Tuesday:
		return tuesday(origin, destination)
	case time.Wednesday:
		return wednesday(origin, destination)
	case time.Thursday:
		return thursday(origin, destination)
	case time.Friday:
		return friday(origin, destination)
	case time.Saturday:
		return saturday(origin, destination)
	case time.Sunday:
		return sunday(origin, destination)
	default:
		return []Flight{}
	}
}

func monday(origin, destination string) []Flight {
	return []Flight{
		{
			FlightNumber: "AA123",
			Origin:       origin,
			Destination:  destination,
			Departure:    time.Now().Add(2 * time.Hour),
			Arrival:      time.Now().Add(4 * time.Hour),
			BasePrice:    199.99,
		},
	}
}

func tuesday(origin, destination string) []Flight {
	return []Flight{
		{
			FlightNumber: "AA124",
			Origin:       origin,
			Destination:  destination,
			Departure:    time.Now().Add(3 * time.Hour),
			Arrival:      time.Now().Add(5 * time.Hour),
			BasePrice:    199.99,
		},
	}
}

func wednesday(origin, destination string) []Flight {
	return []Flight{
		{
			FlightNumber: "AA125",
			Origin:       origin,
			Destination:  destination,
			Departure:    time.Now().Add(4 * time.Hour),
			Arrival:      time.Now().Add(6 * time.Hour),
			BasePrice:    199.93,
		},
	}
}

func thursday(origin, destination string) []Flight {
	return []Flight{
		{
			FlightNumber: "AA126",
			Origin:       origin,
			Destination:  destination,
			Departure:    time.Now().Add(5 * time.Hour),
			Arrival:      time.Now().Add(7 * time.Hour),
			BasePrice:    199.94,
		},
	}
}

func friday(origin, destination string) []Flight {
	return []Flight{
		{
			FlightNumber: "AA127",
			Origin:       origin,
			Destination:  destination,
			Departure:    time.Now().Add(6 * time.Hour),
			Arrival:      time.Now().Add(8 * time.Hour),
			BasePrice:    199.95,
		},
	}
}

func saturday(origin, destination string) []Flight {
	return []Flight{
		{
			FlightNumber: "AA128",
			Origin:       origin,
			Destination:  destination,
			Departure:    time.Now().Add(7 * time.Hour),
			Arrival:      time.Now().Add(9 * time.Hour),
			BasePrice:    199.96,
		},
	}
}

func sunday(origin, destination string) []Flight {
	return []Flight{
		{
			FlightNumber: "AA129",
			Origin:       origin,
			Destination:  destination,
			Departure:    time.Now().Add(8 * time.Hour),
			Arrival:      time.Now().Add(10 * time.Hour),
			BasePrice:    199.97,
		},
	}
}
