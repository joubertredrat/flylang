package latangos

import (
	"time"
)

func flights(origin, destination string, date time.Time) []Flight {
	switch date.Weekday() {
	case time.Sunday:
		return sunday(origin, destination, date)
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
	default:
		return []Flight{}
	}
}

func sunday(origin, destination string, date time.Time) []Flight {
	return []Flight{
		{
			FlightNumber: flightnumber(origin, date),
			Origin:       origin,
			Destination:  destination,
			Departure:    departure(origin, date),
			Arrival:      arrival(origin, date),
			BasePrice:    price(origin, date),
		},
	}
}

func monday(origin, destination string, date time.Time) []Flight {
	return []Flight{
		{
			FlightNumber: flightnumber(origin, date),
			Origin:       origin,
			Destination:  destination,
			Departure:    departure(origin, date),
			Arrival:      arrival(origin, date),
			BasePrice:    price(origin, date),
		},
	}
}

func tuesday(origin, destination string, date time.Time) []Flight {
	return []Flight{
		{
			FlightNumber: flightnumber(origin, date),
			Origin:       origin,
			Destination:  destination,
			Departure:    departure(origin, date),
			Arrival:      arrival(origin, date),
			BasePrice:    price(origin, date),
		},
	}
}

func wednesday(origin, destination string, date time.Time) []Flight {
	return []Flight{
		{
			FlightNumber: flightnumber(origin, date),
			Origin:       origin,
			Destination:  destination,
			Departure:    departure(origin, date),
			Arrival:      arrival(origin, date),
			BasePrice:    price(origin, date),
		},
	}
}

func thursday(origin, destination string, date time.Time) []Flight {
	return []Flight{
		{
			FlightNumber: flightnumber(origin, date),
			Origin:       origin,
			Destination:  destination,
			Departure:    departure(origin, date),
			Arrival:      arrival(origin, date),
			BasePrice:    price(origin, date),
		},
	}
}

func friday(origin, destination string, date time.Time) []Flight {
	return []Flight{
		{
			FlightNumber: flightnumber(origin, date),
			Origin:       origin,
			Destination:  destination,
			Departure:    departure(origin, date),
			Arrival:      arrival(origin, date),
			BasePrice:    price(origin, date),
		},
	}
}

func saturday(origin, destination string, date time.Time) []Flight {
	return []Flight{
		{
			FlightNumber: flightnumber(origin, date),
			Origin:       origin,
			Destination:  destination,
			Departure:    departure(origin, date),
			Arrival:      arrival(origin, date),
			BasePrice:    price(origin, date),
		},
	}
}
