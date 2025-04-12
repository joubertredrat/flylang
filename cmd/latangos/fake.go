package latangos

import (
	"fmt"
	"time"
)

const (
	SCL           = "SCL"
	MIA           = "MIA"
	FLIGHT_PREFIX = "LT"
)

var sufix = map[time.Weekday]map[string]string{
	time.Sunday: {
		SCL: "22",
		MIA: "57",
	},
	time.Monday: {
		SCL: "24",
		MIA: "59",
	},
	time.Tuesday: {
		SCL: "26",
		MIA: "61",
	},
	time.Wednesday: {
		SCL: "28",
		MIA: "63",
	},
	time.Thursday: {
		SCL: "30",
		MIA: "65",
	},
	time.Friday: {
		SCL: "32",
		MIA: "67",
	},
	time.Saturday: {
		SCL: "34",
		MIA: "69",
	},
}

var durationminutes = map[time.Weekday]int{
	time.Sunday:    490,
	time.Monday:    490,
	time.Tuesday:   490,
	time.Wednesday: 490,
	time.Thursday:  490,
	time.Friday:    500,
	time.Saturday:  500,
}

func flightnumber(origin string, date time.Time) string {
	return fmt.Sprintf("%s50%s", FLIGHT_PREFIX, sufix[date.Weekday()][origin])
}

func departure(origin string, date time.Time) time.Time {
	if origin == SCL {
		return time.Date(
			date.Year(), date.Month(), date.Day(),
			9, 30, 0, 0,
			date.Location(),
		)
	}

	return time.Date(
		date.Year(), date.Month(), date.Day(),
		14, 30, 0, 0,
		date.Location(),
	)
}

func arrival(origin string, date time.Time) time.Time {
	return departure(origin, date).Add(time.Duration(durationminutes[date.Weekday()]) * time.Minute)
}
