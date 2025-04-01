package aeronyx

import "time"

func operatingOn(leg FlightLeg, day time.Weekday) bool {
	for _, d := range leg.OperatingDays {
		if d == day {
			return true
		}
	}
	return false
}

func getTimeMultiplier(hour int) float64 {
	switch {
	case hour >= 6 && hour < 12:
		return 1.2
	case hour >= 12 && hour < 18:
		return 1.15
	case hour >= 18 && hour < 23:
		return 0.9
	default:
		return 0.8
	}
}

func CalculatePrice(leg FlightLeg, date time.Time) (float64, bool) {
	if !operatingOn(leg, date.Weekday()) {
		return 0, false
	}

	weekdayMultiplier := map[time.Weekday]float64{
		time.Monday:    1.2,
		time.Tuesday:   0.9,
		time.Wednesday: 0.9,
		time.Thursday:  1.0,
		time.Friday:    1.3,
		time.Saturday:  1.4,
		time.Sunday:    1.4,
	}[date.Weekday()]

	timeMultiplier := getTimeMultiplier(leg.DepartureHour)

	final := leg.BasePrice * weekdayMultiplier * timeMultiplier
	return final, true
}
