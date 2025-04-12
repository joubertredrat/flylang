package copana

import (
	"encoding/xml"
	"time"
)

type (
	Token struct {
		Value  string `xml:"value"`
		Expiry string `xml:"expires_in"`
	}

	Error struct {
		Code    int    `xml:"code"`
		Message string `xml:"message"`
	}

	CustomTime struct {
		time.Time
	}

	Flight struct {
		Code      string     `xml:"flight:code"`
		Departure CustomTime `xml:"date:departure" json:"departure"`
		Arrival   CustomTime `xml:"date:arrival" json:"arrival"`
		Pricing   string     `xml:"amount:pricing"`
		Route     string     `xml:"flight:route"`
	}
)

func (ct CustomTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	formatted := ct.Format("02/01/06 15:04")
	return e.EncodeElement(formatted, start)
}
