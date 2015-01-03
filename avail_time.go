package avail

import (
	"encoding/xml"
	"time"
)

var loc = time.Now().Location()

type AvailTime struct {
	time.Time
}

func (a *AvailTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	d.DecodeElement(&v, &start)

	t, err := time.ParseInLocation("2006-01-02T15:04:05", v, loc)
	if err != nil {
		return err
	}
	*a = AvailTime{t}
	return nil
}
