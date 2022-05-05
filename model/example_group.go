package model

import (
	"encoding/xml"
	"errors"
)

var ErrExampleGroupChildrenUnknownElement = errors.New("example-group have unknown children")

type ExampleGroup struct {
	XMLName              xml.Name            `xml:"example-group"`
	SubID                string              `xml:"subid,attr"`
	Type                 string              `xml:"type,attr"`
	Delimiter            string              `xml:"delimiter,attr"`
	ExampleGroupChildren []ExampleGroupChild `xml:",any"`
}

type ExampleGroupChild struct {
	Type  string
	Value interface{}
}

func (egc *ExampleGroupChild) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	switch start.Name.Local {
	case "example":
		var e Example
		if err := d.DecodeElement(&e, &start); err != nil {
			return err
		}
		egc.Value = e
		egc.Type = start.Name.Local
	case "key":
		var k Key
		if err := d.DecodeElement(&k, &start); err != nil {
			return err
		}
		egc.Value = k
		egc.Type = start.Name.Local
	default:
		return ErrExampleGroupChildrenUnknownElement
	}
	return nil
}
