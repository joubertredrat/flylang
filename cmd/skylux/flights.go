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

func monday(origin, destination string, date time.Time) []Flight {
	return []Flight{
		{
			Callsign: "SL001",
			Routes: []FlightRoute{
				{From: origin, To: "SPJC"},
				{From: "SPJC", To: destination},
			},
			Departure: date.Add(2 * time.Hour),
			Arrival:   date.Add(4 * time.Hour),
			SeatPrice: 150.00,
		},
	}
}

func tuesday(origin, destination string, date time.Time) []Flight {
	return []Flight{
		{
			Callsign: "SL002",
			Routes: []FlightRoute{
				{From: origin, To: "SPJC"},
				{From: "SPJC", To: destination},
			},
			Departure: date.Add(3 * time.Hour),
			Arrival:   date.Add(5 * time.Hour),
			SeatPrice: 160.00,
		},
	}
}

func wednesday(origin, destination string, date time.Time) []Flight {
	return []Flight{
		{
			Callsign: "SL003",
			Routes: []FlightRoute{
				{From: origin, To: "SPJC"},
				{From: "SPJC", To: destination},
			},
			Departure: date.Add(4 * time.Hour),
			Arrival:   date.Add(6 * time.Hour),
			SeatPrice: 170.00,
		},
	}
}

func thursday(origin, destination string, date time.Time) []Flight {
	return []Flight{
		{
			Callsign: "SL004",
			Routes: []FlightRoute{
				{From: origin, To: "SPJC"},
				{From: "SPJC", To: destination},
			},
			Departure: date.Add(5 * time.Hour),
			Arrival:   date.Add(7 * time.Hour),
			SeatPrice: 180.00,
		},
	}
}

func friday(origin, destination string, date time.Time) []Flight {
	return []Flight{
		{
			Callsign: "SL005",
			Routes: []FlightRoute{
				{From: origin, To: "SPJC"},
				{From: "SPJC", To: destination},
			},
			Departure: date.Add(6 * time.Hour),
			Arrival:   date.Add(8 * time.Hour),
			SeatPrice: 190.00,
		},
	}
}

func saturday(origin, destination string, date time.Time) []Flight {
	return []Flight{
		{
			Callsign: "SL006",
			Routes: []FlightRoute{
				{From: origin, To: "SPJC"},
				{From: "SPJC", To: destination},
			},
			Departure: date.Add(7 * time.Hour),
			Arrival:   date.Add(9 * time.Hour),
			SeatPrice: 200.00,
		},
	}
}

func sunday(origin, destination string, date time.Time) []Flight {
	return []Flight{
		{
			Callsign: "SL007",
			Routes: []FlightRoute{
				{From: origin, To: "SPJC"},
				{From: "SPJC", To: destination},
			},
			Departure: date.Add(8 * time.Hour),
			Arrival:   date.Add(10 * time.Hour),
			SeatPrice: 210.00,
		},
	}
}
