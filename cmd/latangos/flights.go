package latangos

import (
	"math"
	"time"
)

var basePrice = float64(735)
var weekMultiplier = map[time.Weekday]float64{
	time.Monday:    1.1,
	time.Tuesday:   1.15,
	time.Wednesday: 1.20,
	time.Thursday:  1.17,
	time.Friday:    0.92,
	time.Saturday:  0.89,
	time.Sunday:    0.94,
}

func price(origin string, date time.Time) float64 {
	wm := weekMultiplier[date.Weekday()]
	dm := float64(1)
	if origin == SCL {
		dm = 1.04
	}

	return math.Round(basePrice*wm*dm*100) / 100
}

func flights(origin, destination string, date time.Time) []Flight {
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
			Arrival:      time.Now().Add(10 * time.Hour),
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
