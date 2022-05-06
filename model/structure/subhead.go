package model

import (
	"encoding/xml"
	"errors"
)

var ErrSubheadChildrenUnknownElement = errors.New("subhead have unknown children")

type Subhead struct {
	Structure
	XMLName         xml.Name       `xml:"subhead"`
	SubID           string         `xml:"subid,attr"`
	Type            string         `xml:"type,attr"`
	Delimiter       string         `xml:"delimiter,attr"`
	SubheadChildren []SubheadChild `xml:",any"`
}

type SubheadChild struct {
	Type  string
	Value interface{}
}

func (sc *SubheadChild) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	switch start.Name.Local {
	case "subheadword":
		var s Subheadword
		if err := d.DecodeElement(&s, &start); err != nil {
			return err
		}
		sc.Value = s
		sc.Type = start.Name.Local
	case "meaning":
		var m Meaning
		if err := d.DecodeElement(&m, &start); err != nil {
			return err
		}
		sc.Value = m
		sc.Type = start.Name.Local
	case "example":
		var e Example
		if err := d.DecodeElement(&e, &start); err != nil {
			return err
		}
		sc.Value = e
		sc.Type = start.Name.Local
	case "snippet":
		var s Snippet
		if err := d.DecodeElement(&s, &start); err != nil {
			return err
		}
		sc.Value = s
		sc.Type = start.Name.Local
	case "column":
		// TODO: column
	case "key":
		var k Key
		if err := d.DecodeElement(&k, &start); err != nil {
			return err
		}
		sc.Value = k
		sc.Type = start.Name.Local
	case "div":
		// TODO: div
	default:
		return ErrSubheadChildrenUnknownElement
	}
	return nil
}
