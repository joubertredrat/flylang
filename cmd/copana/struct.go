package copana

import "time"

type (
	Token struct {
		Token  string `xml:"token"`
		Expiry string `xml:"expires_in"`
	}

	FlightSegment struct {
		IATAFrom string `xml:"iata:departure_code"`
		IATATo   string `xml:"iata:arrival_code"`
	}

	Flight struct {
		Code           string          `xml:"flight:code"`
		Departure      time.Time       `xml:"date:departure"`
		Arrival        time.Time       `xml:"date:arrival"`
		Pricing        int             `xml:"amount:pricing"`
		FlightSegments []FlightSegment `xml:"flight:segments"`
	}
)
