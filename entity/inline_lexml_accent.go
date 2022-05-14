package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrAccentUnknownElement = errors.New("accent have unknown children")

type Accent struct {
	InlineLeXML
	XMLName xml.Name `xml:"accent"`
	Type    string   `xml:"type,attr"`
	Value   InnerXML
}

func (a *Accent) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var aa struct {
		InlineLeXML
		XMLName xml.Name `xml:"accent"`
		Type    string   `xml:"type,attr"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&aa, &start); err != nil {
		return errors.Wrap(ErrAccentUnknownElement, "failed to unmarshal data")
	}
	*a = Accent{
		XMLName: xml.Name{Local: "accent"},
		Type:    aa.Type,
		Value: InnerXML{
			Value: aa.Value,
		},
	}
	return nil
}
