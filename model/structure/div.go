package structure

import (
	"encoding/xml"
	"errors"
)

var ErrDivChildrenUnknownElement = errors.New("div have unknown children")

type Div struct {
	Structure
	XMLName     xml.Name   `xml:"div"`
	SubID       string     `xml:"subid,attr"`
	Type        string     `xml:"type,attr"`
	Level       string     `xml:"level,attr"`
	DivChildren []DivChild `xml:",any"` // (title | key | meaning | example | subhead)*
}

type DivChild struct {
	Type  string
	Value interface{}
}

func (dc *DivChild) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	switch start.Name.Local {
	case "title":
		var t Title
		if err := d.DecodeElement(&t, &start); err != nil {
			return err
		}
		dc.Value = t
		dc.Type = start.Name.Local
	case "key":
		var k Key
		if err := d.DecodeElement(&k, &start); err != nil {
			return err
		}
		dc.Value = k
		dc.Type = start.Name.Local
	case "meaning":
		var m Meaning
		if err := d.DecodeElement(&m, &start); err != nil {
			return err
		}
		dc.Value = m
		dc.Type = start.Name.Local
	case "example":
		var e Example
		if err := d.DecodeElement(&e, &start); err != nil {
			return err
		}
		dc.Value = e
		dc.Type = start.Name.Local
	case "subhead":
		var s Subhead
		if err := d.DecodeElement(&s, &start); err != nil {
			return err
		}
		dc.Value = s
		dc.Type = start.Name.Local
	default:
		return ErrDivChildrenUnknownElement
	}
	return nil
}
