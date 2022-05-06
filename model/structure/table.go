package model

import (
	"encoding/xml"
	"errors"
)

var ErrTableChildrenUnknownElement = errors.New("table have unknown children")

type Table struct {
	Structure
	XMLName    xml.Name  `xml:"table"`
	SubID      string    `xml:"subid,attr"`
	Alt        string    `xml:"alt,attr"`
	Type       string    `xml:"type,attr"`
	TrChildren []TrChild `xml:",any"`
}

type TableChild struct {
	Type  string
	Value interface{}
}

func (tc *TableChild) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	switch start.Name.Local {
	case "caption":
		var c Caption
		if err := d.DecodeElement(&c, &start); err != nil {
			return err
		}
		tc.Value = c
		tc.Type = start.Name.Local
	case "tr":
		var tr Tr
		if err := d.DecodeElement(&tr, &start); err != nil {
			return err
		}
		tc.Value = tr
		tc.Type = start.Name.Local
	default:
		return ErrTableChildrenUnknownElement
	}
	return nil
}
