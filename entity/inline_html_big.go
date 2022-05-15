package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrBigUnknownElement = errors.New("big have unknown children")

type Big struct {
	InlineHTML
	XMLName xml.Name `xml:"big"`
	Type    string   `xml:"type,attr"`
	Class   string   `xml:"class,attr"`
	Value   InnerXML
}

func (b *Big) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var bb struct {
		InlineHTML
		XMLName xml.Name `xml:"big"`
		Type    string   `xml:"type,attr"`
		Class   string   `xml:"class,attr"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&bb, &start); err != nil {
		return errors.Wrap(ErrBigUnknownElement, "failed to unmarshal data")
	}
	*b = Big{
		XMLName: xml.Name{Local: "big"},
		Type:    bb.Type,
		Class:   bb.Class,
		Value: InnerXML{
			Value: bb.Value,
		},
	}
	return nil
}
