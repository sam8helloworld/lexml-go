package model

import (
	"encoding/xml"
	"errors"
)

var ErrIndexChildrenUnknownElement = errors.New("index have unknown children")

type Index struct {
	XMLName       xml.Name     `xml:"index"`
	SubID         string       `xml:"subid,attr"`
	Type          string       `xml:"type,attr"`
	IndexChildren []IndexChild `xml:",any"`
}

type IndexChild struct {
	Type  string
	Value interface{}
}

func (hc *IndexChild) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	switch start.Name.Local {
	case "meaning":
		var m Meaning
		if err := d.DecodeElement(&m, &start); err != nil {
			return err
		}
		hc.Value = m
		hc.Type = start.Name.Local
	case "indexlist":
		var il IndexList
		if err := d.DecodeElement(&il, &start); err != nil {
			return err
		}
		hc.Value = il
		hc.Type = start.Name.Local
	default:
		return ErrIndexChildrenUnknownElement
	}
	return nil
}
