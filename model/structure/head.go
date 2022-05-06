package model

import (
	"encoding/xml"
	"errors"
)

var ErrHeadChildrenUnknownElement = errors.New("head have unknown children")

type Head struct {
	Structure
	XMLName      xml.Name    `xml:"head"`
	HeadChildren []HeadChild `xml:",any"` // Headword or Key
}

type HeadChild struct {
	Type  string
	Value interface{}
}

func (hc *HeadChild) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	switch start.Name.Local {
	case "headword":
		var hh Headword
		if err := d.DecodeElement(&hh, &start); err != nil {
			return err
		}
		hc.Value = hh
		hc.Type = start.Name.Local
	case "key":
		var hh Key
		if err := d.DecodeElement(&hh, &start); err != nil {
			return err
		}
		hc.Value = hh
		hc.Type = start.Name.Local
	default:
		return ErrHeadChildrenUnknownElement
	}
	return nil
}
