package aeronyx

/*
Package aeronyx defines the available itineraries (routes) for the fake airline Aeronyx.

Direct Flights (non-stop):
- GRU → EZE (XV100, XV136)
- GRU → SCL (XV102, XV138)
- GRU → LHR (XV104)
- GRU → MEX (XV106)
- GRU → MIA (XV108, XV140)

Connecting Flights:
- GRU → MIA → ONT (XV140)
- GRU → MEX → ONT (XV152)
- GRU → LHR → WAW (XV154)
- GRU → DXB → DEL (XV156)
- GRU → DXB → HND (XV158)
- GRU → DXB → ICN (XV160)
- GRU → DXB → CAI (XV162)
- GRU → DXB → SYD (XV164)

Each itinerary is defined with one or more legs (FlightLeg), with departure time, base price,
operating days, and duration in minutes. These itineraries are used to simulate realistic
search results in the Aeronyx API.
*/

import "time"

var directItineraries = []Itinerary{
	{
		Legs: []FlightLeg{
			{
				FlightNumber:    "XV100",
				From:            "GRU",
				To:              "EZE",
				BasePrice:       750,
				OperatingDays:   []time.Weekday{time.Sunday, time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday, time.Saturday},
				DepartureHour:   9,
				DepartureTime:   "09:00",
				DurationMinutes: 150,
			},
		},
	},
	{
		Legs: []FlightLeg{
			{
				FlightNumber:    "XV136",
				From:            "GRU",
				To:              "EZE",
				BasePrice:       770,
				OperatingDays:   []time.Weekday{time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday},
				DepartureHour:   18,
				DepartureTime:   "18:00",
				DurationMinutes: 150,
			},
		},
	},
	{
		Legs: []FlightLeg{
			{
				FlightNumber:    "XV102",
				From:            "GRU",
				To:              "SCL",
				BasePrice:       820,
				OperatingDays:   []time.Weekday{time.Sunday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday, time.Saturday},
				DepartureHour:   11,
				DepartureTime:   "11:00",
				DurationMinutes: 210,
			},
		},
	},
	{
		Legs: []FlightLeg{
			{
				FlightNumber:    "XV138",
				From:            "GRU",
				To:              "SCL",
				BasePrice:       840,
				OperatingDays:   []time.Weekday{time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday},
				DepartureHour:   20,
				DepartureTime:   "20:00",
				DurationMinutes: 210,
			},
		},
	},
	{
		Legs: []FlightLeg{
			{
				FlightNumber:    "XV104",
				From:            "GRU",
				To:              "LHR",
				BasePrice:       2700,
				OperatingDays:   []time.Weekday{time.Monday, time.Wednesday},
				DepartureHour:   17,
				DepartureTime:   "17:00",
				DurationMinutes: 660,
			},
		},
	},
	{
		Legs: []FlightLeg{
			{
				FlightNumber:    "XV106",
				From:            "GRU",
				To:              "MEX",
				BasePrice:       1650,
				OperatingDays:   []time.Weekday{time.Tuesday, time.Saturday},
				DepartureHour:   14,
				DepartureTime:   "14:00",
				DurationMinutes: 540,
			},
		},
	},
	{
		Legs: []FlightLeg{
			{
				FlightNumber:    "XV108",
				From:            "GRU",
				To:              "MIA",
				BasePrice:       1600,
				OperatingDays:   []time.Weekday{time.Tuesday, time.Thursday, time.Saturday},
				DepartureHour:   10,
				DepartureTime:   "10:00",
				DurationMinutes: 480,
			},
		},
	},
	{
		Legs: []FlightLeg{
			{
				FlightNumber:    "XV140",
				From:            "GRU",
				To:              "MIA",
				BasePrice:       1650,
				OperatingDays:   []time.Weekday{time.Tuesday, time.Thursday, time.Saturday},
				DepartureHour:   16,
				DepartureTime:   "16:00",
				DurationMinutes: 480,
			},
		},
	},
}

var withConnectionItineraries = []Itinerary{
	{
		Legs: []FlightLeg{
			{
				FlightNumber:    "XV140",
				From:            "GRU",
				To:              "MIA",
				BasePrice:       1600,
				OperatingDays:   []time.Weekday{time.Tuesday, time.Saturday},
				DepartureHour:   10,
				DepartureTime:   "10:00",
				DurationMinutes: 480,
			},
			{
				FlightNumber:    "XV140",
				From:            "MIA",
				To:              "ONT",
				BasePrice:       400,
				OperatingDays:   []time.Weekday{time.Tuesday, time.Saturday},
				DepartureHour:   13,
				DepartureTime:   "13:00",
				DurationMinutes: 120,
			},
		},
	},
	{
		Legs: []FlightLeg{
			{
				FlightNumber:    "XV152",
				From:            "GRU",
				To:              "MEX",
				BasePrice:       1650,
				OperatingDays:   []time.Weekday{time.Thursday, time.Sunday},
				DepartureHour:   14,
				DepartureTime:   "14:00",
				DurationMinutes: 540,
			},
			{
				FlightNumber:    "XV152",
				From:            "MEX",
				To:              "ONT",
				BasePrice:       450,
				OperatingDays:   []time.Weekday{time.Thursday, time.Sunday},
				DepartureHour:   15,
				DepartureTime:   "15:00",
				DurationMinutes: 120,
			},
		},
	},
	{
		Legs: []FlightLeg{
			{
				FlightNumber:    "XV154",
				From:            "GRU",
				To:              "LHR",
				BasePrice:       2700,
				OperatingDays:   []time.Weekday{time.Wednesday},
				DepartureHour:   17,
				DepartureTime:   "17:00",
				DurationMinutes: 660,
			},
			{
				FlightNumber:    "XV154",
				From:            "LHR",
				To:              "WAW",
				BasePrice:       300,
				OperatingDays:   []time.Weekday{time.Wednesday},
				DepartureHour:   21,
				DepartureTime:   "21:00",
				DurationMinutes: 120,
			},
		},
	},
	{
		Legs: []FlightLeg{
			{
				FlightNumber:    "XV156",
				From:            "GRU",
				To:              "DXB",
				BasePrice:       2900,
				OperatingDays:   []time.Weekday{time.Thursday},
				DepartureHour:   22,
				DepartureTime:   "22:00",
				DurationMinutes: 840,
			},
			{
				FlightNumber:    "XV156",
				From:            "DXB",
				To:              "DEL",
				BasePrice:       800,
				OperatingDays:   []time.Weekday{time.Thursday},
				DepartureHour:   2,
				DepartureTime:   "02:00",
				DurationMinutes: 240,
			},
		},
	},
	{
		Legs: []FlightLeg{
			{
				FlightNumber:    "XV158",
				From:            "GRU",
				To:              "DXB",
				BasePrice:       2900,
				OperatingDays:   []time.Weekday{time.Thursday},
				DepartureHour:   22,
				DepartureTime:   "22:00",
				DurationMinutes: 840,
			},
			{
				FlightNumber:    "XV158",
				From:            "DXB",
				To:              "HND",
				BasePrice:       1200,
				OperatingDays:   []time.Weekday{time.Thursday},
				DepartureHour:   2,
				DepartureTime:   "02:00",
				DurationMinutes: 540,
			},
		},
	},
	{
		Legs: []FlightLeg{
			{
				FlightNumber:    "XV160",
				From:            "GRU",
				To:              "DXB",
				BasePrice:       2900,
				OperatingDays:   []time.Weekday{time.Friday},
				DepartureHour:   22,
				DepartureTime:   "22:00",
				DurationMinutes: 840,
			},
			{
				FlightNumber:    "XV161",
				From:            "DXB",
				To:              "ICN",
				BasePrice:       1150,
				OperatingDays:   []time.Weekday{time.Friday},
				DepartureHour:   3,
				DepartureTime:   "03:00",
				DurationMinutes: 510,
			},
		},
	},
	{
		Legs: []FlightLeg{
			{
				FlightNumber:    "XV162",
				From:            "GRU",
				To:              "DXB",
				BasePrice:       2900,
				OperatingDays:   []time.Weekday{time.Monday},
				DepartureHour:   22,
				DepartureTime:   "22:00",
				DurationMinutes: 840,
			},
			{
				FlightNumber:    "XV162",
				From:            "DXB",
				To:              "CAI",
				BasePrice:       900,
				OperatingDays:   []time.Weekday{time.Monday},
				DepartureHour:   23,
				DepartureTime:   "23:00",
				DurationMinutes: 300,
			},
		},
	},
	{
		Legs: []FlightLeg{
			{
				FlightNumber:    "XV164",
				From:            "GRU",
				To:              "DXB",
				BasePrice:       2900,
				OperatingDays:   []time.Weekday{time.Saturday},
				DepartureHour:   8,
				DepartureTime:   "08:00",
				DurationMinutes: 150,
			},
			{
				FlightNumber:    "XV164",
				From:            "DXB",
				To:              "SYD",
				BasePrice:       1000,
				OperatingDays:   []time.Weekday{time.Saturday},
				DepartureHour:   16,
				DepartureTime:   "16:00",
				DurationMinutes: 780,
			},
		},
	},
}

var baseItineraries = append(directItineraries, withConnectionItineraries...)
