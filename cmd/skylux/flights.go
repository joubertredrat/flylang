package skylux

import "time"

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
			Callsign: flightnumber(origin, date),
			Routes: []FlightRoute{
				{From: origin, To: "SPJC"},
				{From: "SPJC", To: destination},
			},
			Departure: departure(origin, date),
			Arrival:   arrival(origin, date),
			SeatPrice: price(origin, date),
		},
	}
}

func monday(origin, destination string, date time.Time) []Flight {
	return []Flight{}
}

func tuesday(origin, destination string, date time.Time) []Flight {
	return []Flight{
		{
			Callsign: flightnumber(origin, date),
			Routes: []FlightRoute{
				{From: origin, To: "SPJC"},
				{From: "SPJC", To: destination},
			},
			Departure: departure(origin, date),
			Arrival:   arrival(origin, date),
			SeatPrice: price(origin, date),
		},
	}
}

func wednesday(origin, destination string, date time.Time) []Flight {
	return []Flight{
		{
			Callsign: flightnumber(origin, date),
			Routes: []FlightRoute{
				{From: origin, To: "SPJC"},
				{From: "SPJC", To: destination},
			},
			Departure: departure(origin, date),
			Arrival:   arrival(origin, date),
			SeatPrice: price(origin, date),
		},
	}
}

func thursday(origin, destination string, date time.Time) []Flight {
	return []Flight{
		{
			Callsign: flightnumber(origin, date),
			Routes: []FlightRoute{
				{From: origin, To: "SPJC"},
				{From: "SPJC", To: destination},
			},
			Departure: departure(origin, date),
			Arrival:   arrival(origin, date),
			SeatPrice: price(origin, date),
		},
	}
}

func friday(origin, destination string, date time.Time) []Flight {
	return []Flight{
		{
			Callsign: flightnumber(origin, date),
			Routes: []FlightRoute{
				{From: origin, To: "SPJC"},
				{From: "SPJC", To: destination},
			},
			Departure: departure(origin, date),
			Arrival:   arrival(origin, date),
			SeatPrice: price(origin, date),
		},
	}
}

func saturday(origin, destination string, date time.Time) []Flight {
	return []Flight{
		{
			Callsign: flightnumber(origin, date),
			Routes: []FlightRoute{
				{From: origin, To: "SPJC"},
				{From: "SPJC", To: destination},
			},
			Departure: departure(origin, date),
			Arrival:   arrival(origin, date),
			SeatPrice: price(origin, date),
		},
	}
}
