package aeronyx

import (
	"fmt"
	"time"
)

func generateReturnRoutes(routes []Itinerary) []Itinerary {
	var returns []Itinerary
	flightNum := 101

	for _, route := range routes {
		if len(route.Legs) == 0 {
			continue
		}

		reversed := make([]FlightLeg, 0, len(route.Legs))
		for i := len(route.Legs) - 1; i >= 0; i-- {
			leg := route.Legs[i]
			reversed = append(reversed, FlightLeg{
				FlightNumber:    "XV" + formatFlightNumber(flightNum),
				From:            leg.To,
				To:              leg.From,
				BasePrice:       leg.BasePrice,
				OperatingDays:   leg.OperatingDays,
				DepartureHour:   leg.DepartureHour,
				DepartureTime:   leg.DepartureTime,
				DurationMinutes: leg.DurationMinutes,
			})
			flightNum += 2
		}

		returns = append(returns, Itinerary{Legs: reversed})
	}

	return returns
}

func formatFlightNumber(n int) string {
	return fmt.Sprintf("%03d", n)
}

var itineraries = append(baseItineraries, generateReturnRoutes(baseItineraries)...)

func SearchFlights(from, to string, date time.Time) []FlightResponse {
	results := []FlightResponse{}

	for _, itinerary := range itineraries {
		if len(itinerary.Legs) == 0 {
			continue
		}

		origin := itinerary.Legs[0].From
		destination := itinerary.Legs[len(itinerary.Legs)-1].To

		if origin != from || destination != to {
			continue
		}

		totalPrice := 0.0
		valid := true
		totalDuration := 0
		for _, leg := range itinerary.Legs {
			price, ok := CalculatePrice(leg, date)
			if !ok {
				valid = false
				break
			}
			totalPrice += price
			totalDuration += leg.DurationMinutes
		}

		if valid {
			firstLeg := itinerary.Legs[0]

			results = append(results, FlightResponse{
				FlightNumber:         firstLeg.FlightNumber,
				From:                 origin,
				To:                   destination,
				TotalPrice:           fmt.Sprintf("%.2f", totalPrice),
				DepartureTime:        firstLeg.DepartureTime,
				TotalDurationMinutes: totalDuration,
				NumberOfConnections:  len(itinerary.Legs) - 1,
			})
		}
	}

	return results
}
