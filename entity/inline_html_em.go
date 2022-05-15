package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrEmUnknownElement = errors.New("em have unknown children")

type SmallEM struct {
	InlineHTML
	XMLName xml.Name `xml:"em"`
	Type    string   `xml:"type,attr"`
	Class   string   `xml:"class,attr"`
	Value   InnerXML
}

type LargeEM struct {
	InlineHTML
	XMLName xml.Name `xml:"EM"`
	Type    string   `xml:"type,attr"`
	Class   string   `xml:"class,attr"`
	Value   InnerXML
}

func (e *SmallEM) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var ee struct {
		InlineHTML
		XMLName xml.Name `xml:"em"`
		Type    string   `xml:"type,attr"`
		Class   string   `xml:"class,attr"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&ee, &start); err != nil {
		return errors.Wrap(ErrEmUnknownElement, "failed to unmarshal data")
	}
	*e = SmallEM{
		XMLName: xml.Name{Local: "em"},
		Type:    ee.Type,
		Class:   ee.Class,
		Value: InnerXML{
			Value: ee.Value,
		},
	}
	return nil
}

func (e *LargeEM) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var ee struct {
		InlineHTML
		XMLName xml.Name `xml:"EM"`
		Type    string   `xml:"type,attr"`
		Class   string   `xml:"class,attr"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&ee, &start); err != nil {
		return errors.Wrap(ErrEmUnknownElement, "failed to unmarshal data")
	}
	*e = LargeEM{
		XMLName: xml.Name{Local: "EM"},
		Type:    ee.Type,
		Class:   ee.Class,
		Value: InnerXML{
			Value: ee.Value,
		},
	}
	return nil
}
