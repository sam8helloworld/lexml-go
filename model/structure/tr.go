package structure

import (
	"encoding/xml"
	"errors"
)

var ErrTrChildrenUnknownElement = errors.New("tr have unknown children")

type Tr struct {
	Structure
	XMLName    xml.Name  `xml:"tr"`
	TrChildren []TrChild `xml:",any"` // (th | td)*
}

type TrChild struct {
	Type  string
	Value interface{}
}

func (tc *TrChild) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	switch start.Name.Local {
	case "th":
		var th Th
		if err := d.DecodeElement(&th, &start); err != nil {
			return err
		}
		tc.Value = th
		tc.Type = start.Name.Local
	case "td":
		var td Td
		if err := d.DecodeElement(&td, &start); err != nil {
			return err
		}
		tc.Value = td
		tc.Type = start.Name.Local
	default:
		return ErrTrChildrenUnknownElement
	}
	return nil
}
