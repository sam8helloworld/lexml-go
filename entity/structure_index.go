package entity

import (
	"encoding/xml"
	"errors"
)

var ErrIndexChildrenUnknownElement = errors.New("index have unknown children")

type Index struct {
	Structure
	XMLName       xml.Name     `xml:"index"`
	SubID         string       `xml:"subid,attr"`
	Type          string       `xml:"type,attr"`
	IndexChildren []IndexChild `xml:",any"` // (meaning | indexlist)*
}

type IndexChild struct {
	Type  string
	Value interface{}
}

func (ic *IndexChild) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	switch start.Name.Local {
	case "meaning":
		var m Meaning
		if err := d.DecodeElement(&m, &start); err != nil {
			return err
		}
		ic.Value = m
		ic.Type = start.Name.Local
	case "indexlist":
		var il IndexList
		if err := d.DecodeElement(&il, &start); err != nil {
			return err
		}
		ic.Value = il
		ic.Type = start.Name.Local
	default:
		return ErrIndexChildrenUnknownElement
	}
	return nil
}
