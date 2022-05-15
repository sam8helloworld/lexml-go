package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrBUnknownElement = errors.New("b have unknown children")

type SmallB struct {
	InlineHTML
	XMLName xml.Name `xml:"b"`
	Type    string   `xml:"type,attr"`
	Class   string   `xml:"class,attr"`
	Value   InnerXML
}

type LargeB struct {
	InlineHTML
	XMLName xml.Name `xml:"B"`
	Type    string   `xml:"type,attr"`
	Class   string   `xml:"class,attr"`
	Value   InnerXML
}

func (b *SmallB) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var bb struct {
		InlineHTML
		XMLName xml.Name `xml:"b"`
		Type    string   `xml:"type,attr"`
		Class   string   `xml:"class,attr"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&bb, &start); err != nil {
		return errors.Wrap(ErrBUnknownElement, "failed to unmarshal data")
	}
	*b = SmallB{
		XMLName: xml.Name{Local: "b"},
		Type:    bb.Type,
		Class:   bb.Class,
		Value: InnerXML{
			Value: bb.Value,
		},
	}
	return nil
}

func (b *LargeB) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var bb struct {
		InlineHTML
		XMLName xml.Name `xml:"B"`
		Type    string   `xml:"type,attr"`
		Class   string   `xml:"class,attr"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&bb, &start); err != nil {
		return errors.Wrap(ErrBUnknownElement, "failed to unmarshal data")
	}
	*b = LargeB{
		XMLName: xml.Name{Local: "B"},
		Type:    bb.Type,
		Class:   bb.Class,
		Value: InnerXML{
			Value: bb.Value,
		},
	}
	return nil
}
