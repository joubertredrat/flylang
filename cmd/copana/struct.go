package copana

import "time"

type (
	Token struct {
		Value  string `xml:"value"`
		Expiry string `xml:"expires_in"`
	}

	Flight struct {
		Code      string    `xml:"flight:code"`
		Departure time.Time `xml:"date:departure"`
		Arrival   time.Time `xml:"date:arrival"`
		Pricing   int       `xml:"amount:pricing"`
		Route     string    `xml:"flight:route"`
	}
)
