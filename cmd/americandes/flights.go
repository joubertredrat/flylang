package americandes

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
			ID:                "AM123",
			IATACodeFrom:      origin,
			IATACodeTo:        destination,
			DepartureTime:     time.Now().Add(2 * time.Hour),
			DurationInMinutes: 120,
			Amount:            250,
		},
	}
}

func tuesday(origin, destination string) []Flight {
	return []Flight{
		{
			ID:                "AM124",
			IATACodeFrom:      origin,
			IATACodeTo:        destination,
			DepartureTime:     time.Now().Add(3 * time.Hour),
			DurationInMinutes: 150,
			Amount:            260,
		},
	}
}

func wednesday(origin, destination string) []Flight {
	return []Flight{
		{
			ID:                "AM125",
			IATACodeFrom:      origin,
			IATACodeTo:        destination,
			DepartureTime:     time.Now().Add(4 * time.Hour),
			DurationInMinutes: 180,
			Amount:            270,
		},
	}
}

func thursday(origin, destination string) []Flight {
	return []Flight{
		{
			ID:                "AM126",
			IATACodeFrom:      origin,
			IATACodeTo:        destination,
			DepartureTime:     time.Now().Add(5 * time.Hour),
			DurationInMinutes: 200,
			Amount:            280,
		},
	}
}

func friday(origin, destination string) []Flight {
	return []Flight{
		{
			ID:                "AM127",
			IATACodeFrom:      origin,
			IATACodeTo:        destination,
			DepartureTime:     time.Now().Add(6 * time.Hour),
			DurationInMinutes: 220,
			Amount:            290,
		},
	}
}

func saturday(origin, destination string) []Flight {
	return []Flight{
		{
			ID:                "AM128",
			IATACodeFrom:      origin,
			IATACodeTo:        destination,
			DepartureTime:     time.Now().Add(7 * time.Hour),
			DurationInMinutes: 240,
			Amount:            300,
		},
	}
}

func sunday(origin, destination string) []Flight {
	return []Flight{
		{
			ID:                "AM129",
			IATACodeFrom:      origin,
			IATACodeTo:        destination,
			DepartureTime:     time.Now().Add(8 * time.Hour),
			DurationInMinutes: 260,
			Amount:            310,
		},
	}
}
