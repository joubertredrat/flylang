package americandes

import (
	"fmt"
	"math"
	"time"
)

const (
	SCL           = "SCL"
	MIA           = "MIA"
	FLIGHT_PREFIX = "ADS"
)

var basePrice = float64(940)
var weekMultiplier = map[time.Weekday]float64{
	time.Monday:    0.9,
	time.Tuesday:   0.85,
	time.Wednesday: 0.80,
	time.Thursday:  0.85,
	time.Friday:    1.10,
	time.Saturday:  1.15,
	time.Sunday:    1.20,
}

var sufix = map[time.Weekday]map[string]string{
	time.Monday: {
		SCL: "127",
		MIA: "114",
	},
	time.Tuesday: {
		SCL: "128",
		MIA: "115",
	},
	time.Wednesday: {
		SCL: "150",
		MIA: "116",
	},
	time.Thursday: {
		SCL: "155",
		MIA: "170",
	},
	time.Friday: {
		SCL: "157",
		MIA: "172",
	},
	time.Saturday: {
		SCL: "168",
		MIA: "174",
	},
}

var duration = map[time.Weekday]int{
	time.Monday:    480,
	time.Tuesday:   480,
	time.Wednesday: 480,
	time.Thursday:  480,
	time.Friday:    490,
	time.Saturday:  490,
}

func price(origin string, date time.Time) int {
	wm := weekMultiplier[date.Weekday()]
	dm := float64(1)
	if origin == SCL {
		dm = 1.02
	}

	p := math.Round(basePrice*wm*dm*100) / 100
	return int(p * 100)
}

func flightnumber(origin string, date time.Time) string {
	return fmt.Sprintf("%s%s", FLIGHT_PREFIX, sufix[date.Weekday()][origin])
}

func departure(origin string, date time.Time) time.Time {
	if origin == SCL {
		return time.Date(
			date.Year(), date.Month(), date.Day(),
			18, 45, 0, 0,
			date.Location(),
		)
	}

	return time.Date(
		date.Year(), date.Month(), date.Day(),
		12, 25, 0, 0,
		date.Location(),
	)
}

func durationminutes(date time.Time) int {
	return duration[date.Weekday()]
}
