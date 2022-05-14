package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrPhAUnknownElement = errors.New("pha have unknown children")

type PhA struct {
	InlineLeXML
	XMLName xml.Name `xml:"pha"`
	Type    string   `xml:"type,attr"`
	Value   InnerXML
}

func (p *PhA) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var pp struct {
		InlineLeXML
		XMLName xml.Name `xml:"pha"`
		Type    string   `xml:"type,attr"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&pp, &start); err != nil {
		return errors.Wrap(ErrPhAUnknownElement, "failed to unmarshal data")
	}
	*p = PhA{
		XMLName: xml.Name{Local: "pha"},
		Type:    pp.Type,
		Value: InnerXML{
			Value: pp.Value,
		},
	}
	return nil
}
