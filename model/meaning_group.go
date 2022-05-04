package model

import (
	"encoding/xml"
	"errors"
)

var ErrMeaningGroupChildrenUnknownElement = errors.New("meaning-group have unknown children")

type MeaningGroup struct {
	XMLName              xml.Name            `xml:"meaning-group"`
	SubID                string              `xml:"subid,attr"`
	Type                 string              `xml:"type,attr"`
	Level                string              `xml:"level,attr"`
	No                   string              `xml:"no,attr"`
	MeaningGroupChildren []MeaningGroupChild `xml:",any"`
}

type MeaningGroupChild struct {
	Type  string
	Value interface{}
}

func (mgc *MeaningGroupChild) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	switch start.Name.Local {
	case "meaning":
		var m Meaning
		if err := d.DecodeElement(&m, &start); err != nil {
			return err
		}
		mgc.Value = m
		mgc.Type = start.Name.Local
	case "example":
		var e Example
		if err := d.DecodeElement(&e, &start); err != nil {
			return err
		}
		mgc.Value = e
		mgc.Type = start.Name.Local
	case "subhead":
		// TODO: subhead
	case "column":
		// TODO: column
	case "key":
		var k Key
		if err := d.DecodeElement(&k, &start); err != nil {
			return err
		}
		mgc.Value = k
		mgc.Type = start.Name.Local
	case "div":
		// TODO: div
	default:
		return ErrMeaningGroupChildrenUnknownElement
	}
	return nil
}
