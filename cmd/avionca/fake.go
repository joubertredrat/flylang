package avionca

import (
	"fmt"
	"math"
	"time"
)

const (
	SCL           = "SCL"
	MIA           = "MIA"
	FLIGHT_PREFIX = "VV"
)

var baseprice = float64(400)
var weekmultiplier = map[time.Weekday]float64{
	time.Sunday:    1.3,
	time.Monday:    0.975,
	time.Tuesday:   0.925,
	time.Wednesday: 0.95,
	time.Thursday:  1.1,
	time.Saturday:  1.2,
}

var sufix = map[time.Weekday]map[string]string{
	time.Sunday: {
		SCL: "3337",
		MIA: "3338",
	},
	time.Monday: {
		SCL: "3347",
		MIA: "3348",
	},
	time.Tuesday: {
		SCL: "3367",
		MIA: "3368",
	},
	time.Wednesday: {
		SCL: "3377",
		MIA: "3378",
	},
	time.Thursday: {
		SCL: "3417",
		MIA: "3418",
	},
	time.Saturday: {
		SCL: "3477",
		MIA: "3478",
	},
}

var durationminutes = map[time.Weekday]int{
	time.Sunday:    1240,
	time.Monday:    1160,
	time.Tuesday:   1160,
	time.Wednesday: 1160,
	time.Thursday:  1160,
	time.Saturday:  1240,
}

func price(origin string, date time.Time) float64 {
	wm := weekmultiplier[date.Weekday()]
	dm := float64(1)
	if origin == SCL {
		dm = 1.07
	}

	return math.Round(baseprice*wm*dm*100) / 100
}

func flightnumber(origin string, date time.Time) string {
	return fmt.Sprintf("%s%s", FLIGHT_PREFIX, sufix[date.Weekday()][origin])
}

func departure(origin string, date time.Time) string {
	var departureTime time.Time
	if origin == SCL {
		departureTime = time.Date(
			date.Year(), date.Month(), date.Day(),
			9, 30, 0, 0,
			date.Location(),
		)
	} else {
		departureTime = time.Date(
			date.Year(), date.Month(), date.Day(),
			14, 30, 0, 0,
			date.Location(),
		)
	}

	return departureTime.Format("2006-01-02T15:04:05Z")
}

func arrival(date time.Time) int {
	return durationminutes[date.Weekday()]
}
