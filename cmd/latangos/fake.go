package latangos

import (
	"fmt"
	"math"
	"time"
)

const (
	SCL           = "SCL"
	MIA           = "MIA"
	FLIGHT_PREFIX = "LT"
)

var baseprice = float64(735)
var weekmultiplier = map[time.Weekday]float64{
	time.Sunday:    0.94,
	time.Monday:    1.1,
	time.Tuesday:   1.15,
	time.Wednesday: 1.20,
	time.Thursday:  1.17,
	time.Friday:    0.92,
	time.Saturday:  0.89,
}

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

func price(origin string, date time.Time) float64 {
	wm := weekmultiplier[date.Weekday()]
	dm := float64(1)
	if origin == SCL {
		dm = 1.04
	}

	return math.Round(baseprice*wm*dm*100) / 100
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
