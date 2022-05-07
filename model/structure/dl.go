package structure

import (
	"encoding/xml"
	"errors"
)

var ErrDlChildrenUnknownElement = errors.New("dl have unknown children")

type Dl struct {
	Structure
	XMLName    xml.Name  `xml:"dl"`
	SubID      string    `xml:"subid,attr"`
	Type       string    `xml:"type,attr"`
	DlChildren []DlChild `xml:",any"`
}

type DlChild struct {
	Type  string
	Value interface{}
}

func (dc *DlChild) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	switch start.Name.Local {
	case "dt":
		var dt Dt
		if err := d.DecodeElement(&dt, &start); err != nil {
			return err
		}
		dc.Value = dt
		dc.Type = start.Name.Local
	case "dd":
		var dd Dd
		if err := d.DecodeElement(&dd, &start); err != nil {
			return err
		}
		dc.Value = dd
		dc.Type = start.Name.Local
	default:
		return ErrDlChildrenUnknownElement
	}
	return nil
}
