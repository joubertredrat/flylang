package copana

import (
	"fmt"
	"time"
)

func flights(origin, destination string, date time.Time) []Flight {
	switch date.Weekday() {
	case time.Monday:
		return monday(origin, destination, date)
	case time.Tuesday:
		return []Flight{}
	case time.Wednesday:
		return []Flight{}
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

func monday(origin, destination string, date time.Time) []Flight {
	return []Flight{
		{
			Code:      flightnumber(origin, date),
			Departure: departure(origin, date),
			Arrival:   arrival(origin, date),
			Pricing:   price(origin, date),
			Route:     fmt.Sprintf("%s->PTY->%s", origin, destination),
		},
	}
}

func thursday(origin, destination string, date time.Time) []Flight {
	return []Flight{
		{
			Code:      flightnumber(origin, date),
			Departure: departure(origin, date),
			Arrival:   arrival(origin, date),
			Pricing:   price(origin, date),
			Route:     fmt.Sprintf("%s->PTY->%s", origin, destination),
		},
	}
}

func friday(origin, destination string, date time.Time) []Flight {
	return []Flight{
		{
			Code:      flightnumber(origin, date),
			Departure: departure(origin, date),
			Arrival:   arrival(origin, date),
			Pricing:   price(origin, date),
			Route:     fmt.Sprintf("%s->PTY->%s", origin, destination),
		},
	}
}

func saturday(origin, destination string, date time.Time) []Flight {
	return []Flight{
		{
			Code:      flightnumber(origin, date),
			Departure: departure(origin, date),
			Arrival:   arrival(origin, date),
			Pricing:   price(origin, date),
			Route:     fmt.Sprintf("%s->PTY->%s", origin, destination),
		},
	}
}

func sunday(origin, destination string, date time.Time) []Flight {
	return []Flight{
		{
			Code:      flightnumber(origin, date),
			Departure: departure(origin, date),
			Arrival:   arrival(origin, date),
			Pricing:   price(origin, date),
			Route:     fmt.Sprintf("%s->PTY->%s", origin, destination),
		},
	}
}
