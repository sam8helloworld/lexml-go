package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrUUnknownElement = errors.New("u have unknown children")

type SmallU struct {
	InlineHTML
	XMLName xml.Name `xml:"u"`
	Type    string   `xml:"type,attr"`
	Class   string   `xml:"class,attr"`
	Value   InnerXML
}

type LargeU struct {
	InlineHTML
	XMLName xml.Name `xml:"U"`
	Type    string   `xml:"type,attr"`
	Class   string   `xml:"class,attr"`
	Value   InnerXML
}

func (s *SmallU) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var ss struct {
		InlineHTML
		XMLName xml.Name `xml:"u"`
		Type    string   `xml:"type,attr"`
		Class   string   `xml:"class,attr"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&ss, &start); err != nil {
		return errors.Wrap(ErrUUnknownElement, "failed to unmarshal data")
	}
	*s = SmallU{
		XMLName: xml.Name{Local: "u"},
		Type:    ss.Type,
		Class:   ss.Class,
		Value: InnerXML{
			Value: ss.Value,
		},
	}
	return nil
}

func (s *LargeU) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var ss struct {
		InlineHTML
		XMLName xml.Name `xml:"U"`
		Type    string   `xml:"type,attr"`
		Class   string   `xml:"class,attr"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&ss, &start); err != nil {
		return errors.Wrap(ErrUUnknownElement, "failed to unmarshal data")
	}
	*s = LargeU{
		XMLName: xml.Name{Local: "U"},
		Type:    ss.Type,
		Class:   ss.Class,
		Value: InnerXML{
			Value: ss.Value,
		},
	}
	return nil
}
