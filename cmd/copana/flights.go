package copana

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
			Code:      "CP001",
			Departure: date.Add(2 * time.Hour),
			Arrival:   date.Add(4 * time.Hour),
			Pricing:   300,
			FlightSegments: []FlightSegment{
				{
					IATAFrom: origin,
					IATATo:   destination,
				},
			},
		},
	}
}

func tuesday(origin, destination string, date time.Time) []Flight {
	return []Flight{
		{
			Code:      "CP002",
			Departure: date.Add(3 * time.Hour),
			Arrival:   date.Add(5 * time.Hour),
			Pricing:   320,
			FlightSegments: []FlightSegment{
				{
					IATAFrom: origin,
					IATATo:   destination,
				},
			},
		},
	}
}

func wednesday(origin, destination string, date time.Time) []Flight {
	return []Flight{
		{
			Code:      "CP003",
			Departure: date.Add(4 * time.Hour),
			Arrival:   date.Add(6 * time.Hour),
			Pricing:   340,
			FlightSegments: []FlightSegment{
				{
					IATAFrom: origin,
					IATATo:   destination,
				},
			},
		},
	}
}

func thursday(origin, destination string, date time.Time) []Flight {
	return []Flight{
		{
			Code:      "CP004",
			Departure: date.Add(5 * time.Hour),
			Arrival:   date.Add(7 * time.Hour),
			Pricing:   360,
			FlightSegments: []FlightSegment{
				{
					IATAFrom: origin,
					IATATo:   destination,
				},
			},
		},
	}
}

func friday(origin, destination string, date time.Time) []Flight {
	return []Flight{
		{
			Code:      "CP005",
			Departure: date.Add(6 * time.Hour),
			Arrival:   date.Add(8 * time.Hour),
			Pricing:   380,
			FlightSegments: []FlightSegment{
				{
					IATAFrom: origin,
					IATATo:   destination,
				},
			},
		},
	}
}

func saturday(origin, destination string, date time.Time) []Flight {
	return []Flight{
		{
			Code:      "CP006",
			Departure: date.Add(7 * time.Hour),
			Arrival:   date.Add(9 * time.Hour),
			Pricing:   400,
			FlightSegments: []FlightSegment{
				{
					IATAFrom: origin,
					IATATo:   destination,
				},
			},
		},
	}
}

func sunday(origin, destination string, date time.Time) []Flight {
	return []Flight{
		{
			Code:      "CP007",
			Departure: date.Add(8 * time.Hour),
			Arrival:   date.Add(10 * time.Hour),
			Pricing:   420,
			FlightSegments: []FlightSegment{
				{
					IATAFrom: origin,
					IATATo:   destination,
				},
			},
		},
	}
}
