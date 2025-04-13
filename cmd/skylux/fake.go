package skylux

import (
	"fmt"
	"math"
	"time"
)

const (
	SCEL          = "SCEL"
	KMIA          = "KMIA"
	FLIGHT_PREFIX = "SKL"
)

var baseprice = float64(445)
var weekmultiplier = map[time.Weekday]float64{
	time.Sunday:    1.112,
	time.Tuesday:   0.894,
	time.Wednesday: 0.91,
	time.Thursday:  1.179,
	time.Friday:    1.345,
	time.Saturday:  1.371,
}

var sufix = map[time.Weekday]map[string]string{
	time.Sunday: {
		SCEL: "626",
		KMIA: "627",
	},
	time.Tuesday: {
		SCEL: "628",
		KMIA: "629",
	},
	time.Wednesday: {
		SCEL: "630",
		KMIA: "631",
	},
	time.Thursday: {
		SCEL: "632",
		KMIA: "633",
	},
	time.Friday: {
		SCEL: "634",
		KMIA: "635",
	},
	time.Saturday: {
		SCEL: "636",
		KMIA: "637",
	},
}

var durationminutes = map[time.Weekday]int{
	time.Sunday:    510,
	time.Tuesday:   510,
	time.Wednesday: 510,
	time.Thursday:  510,
	time.Friday:    510,
	time.Saturday:  510,
}

func price(origin string, date time.Time) float64 {
	wm := weekmultiplier[date.Weekday()]
	dm := float64(1)
	if origin == SCEL {
		dm = 1.01
	}

	return math.Round(baseprice*wm*dm*100) / 100
}

func flightnumber(origin string, date time.Time) string {
	return fmt.Sprintf("%s0%s", FLIGHT_PREFIX, sufix[date.Weekday()][origin])
}

func departure(origin string, date time.Time) time.Time {
	if origin == SCEL {
		return time.Date(
			date.Year(), date.Month(), date.Day(),
			12, 10, 0, 0,
			date.Location(),
		)
	}

	return time.Date(
		date.Year(), date.Month(), date.Day(),
		12, 50, 0, 0,
		date.Location(),
	)
}

func arrival(origin string, date time.Time) time.Time {
	return departure(origin, date).Add(time.Duration(durationminutes[date.Weekday()]) * time.Minute)
}
