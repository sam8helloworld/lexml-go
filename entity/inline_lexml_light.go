package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrLightUnknownElement = errors.New("light have unknown children")

type Light struct {
	InlineLeXML
	XMLName xml.Name `xml:"light"`
	Value   InnerXML
}

func (l *Light) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var ll struct {
		InlineLeXML
		XMLName xml.Name `xml:"light"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&ll, &start); err != nil {
		return errors.Wrap(ErrLightUnknownElement, "failed to unmarshal data")
	}
	*l = Light{
		XMLName: xml.Name{Local: "light"},
		Value: InnerXML{
			Value: ll.Value,
		},
	}
	return nil
}
