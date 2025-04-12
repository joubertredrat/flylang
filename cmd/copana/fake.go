package copana

import (
	"fmt"
	"math"
	"time"
)

const (
	SCL           = "SCL"
	MIA           = "MIA"
	FLIGHT_PREFIX = "CP"
)

var baseprice = float64(592)
var weekmultiplier = map[time.Weekday]float64{
	time.Sunday:   0.9,
	time.Monday:   1,
	time.Thursday: 0.95,
	time.Friday:   1.1,
	time.Saturday: 1.15,
}

var sufix = map[time.Weekday]map[string]string{
	time.Sunday: {
		SCL: "20",
		MIA: "25",
	},
	time.Monday: {
		SCL: "21",
		MIA: "26",
	},
	time.Thursday: {
		SCL: "22",
		MIA: "27",
	},
	time.Friday: {
		SCL: "23",
		MIA: "28",
	},
	time.Saturday: {
		SCL: "24",
		MIA: "29",
	},
}

var durationminutes = map[time.Weekday]int{
	time.Sunday:   735,
	time.Monday:   715,
	time.Thursday: 715,
	time.Friday:   735,
	time.Saturday: 735,
}

func price(origin string, date time.Time) string {
	wm := weekmultiplier[date.Weekday()]
	dm := float64(1)
	if origin == SCL {
		dm = 1.17
	}

	price := math.Round(baseprice*wm*dm*100000) / 100000
	return fmt.Sprintf("%.5f", price)
}

func flightnumber(origin string, date time.Time) string {
	return fmt.Sprintf("%s%s", FLIGHT_PREFIX, sufix[date.Weekday()][origin])
}

func departure(origin string, date time.Time) CustomTime {
	if origin == SCL {
		return CustomTime{time.Date(
			date.Year(), date.Month(), date.Day(),
			22, 10, 0, 0,
			date.Location(),
		)}
	}

	return CustomTime{time.Date(
		date.Year(), date.Month(), date.Day(),
		05, 30, 0, 0,
		date.Location(),
	)}
}

func arrival(origin string, date time.Time) CustomTime {
	return CustomTime{
		departure(origin, date).Add(time.Duration(durationminutes[date.Weekday()]) * time.Minute),
	}
}
