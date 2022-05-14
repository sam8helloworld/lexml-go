package entity

import (
	"encoding/xml"
	"errors"
)

var ErrReplaceChildrenUnknownElement = errors.New("head have unknown children")

type Replace struct {
	Structure
	XMLName         xml.Name       `xml:"replace"`
	SubId           string         `xml:"subid,attr"`
	Type            string         `xml:"type,attr"`
	Src             string         `xml:"src,attr"`
	ReplaceChildren []ReplaceChild `xml:",any"` //  (meaning | example | subhead)*
}

type ReplaceChild struct {
	Type  string
	Value interface{}
}

func (rc *ReplaceChild) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	switch start.Name.Local {
	case "meaning":
		var m Meaning
		if err := d.DecodeElement(&m, &start); err != nil {
			return err
		}
		rc.Value = m
		rc.Type = start.Name.Local
	case "example":
		var e Example
		if err := d.DecodeElement(&e, &start); err != nil {
			return err
		}
		rc.Value = e
		rc.Type = start.Name.Local
	case "subhead":
		var s Subhead
		if err := d.DecodeElement(&s, &start); err != nil {
			return err
		}
		rc.Value = s
		rc.Type = start.Name.Local
	default:
		return ErrReplaceChildrenUnknownElement
	}
	return nil
}
