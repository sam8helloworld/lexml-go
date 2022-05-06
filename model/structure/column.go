package model

import (
	"encoding/xml"
	"errors"
)

var ErrColumnChildrenUnknownElement = errors.New("column have unknown children")

type Column struct {
	Structure
	XMLName        xml.Name      `xml:"column"`
	SubID          string        `xml:"subid,attr"`
	Type           string        `xml:"type,attr"`
	ColumnChildren []ColumnChild `xml:",any"`
}

type ColumnChild struct {
	Type  string
	Value interface{}
}

func (cc *ColumnChild) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	switch start.Name.Local {
	case "title":
		var t Title
		if err := d.DecodeElement(&t, &start); err != nil {
			return err
		}
		cc.Value = t
		cc.Type = start.Name.Local
	case "key":
		var k Key
		if err := d.DecodeElement(&k, &start); err != nil {
			return err
		}
		cc.Value = k
		cc.Type = start.Name.Local
	case "meaning":
		var m Meaning
		if err := d.DecodeElement(&m, &start); err != nil {
			return err
		}
		cc.Value = m
		cc.Type = start.Name.Local
	case "example":
		var e Example
		if err := d.DecodeElement(&e, &start); err != nil {
			return err
		}
		cc.Value = e
		cc.Type = start.Name.Local
	case "subhead":
		var s Subhead
		if err := d.DecodeElement(&s, &start); err != nil {
			return err
		}
		cc.Value = s
		cc.Type = start.Name.Local
	default:
		return ErrColumnChildrenUnknownElement
	}
	return nil
}
