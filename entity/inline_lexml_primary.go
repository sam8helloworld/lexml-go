package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrPrimaryUnknownElement = errors.New("primary have unknown children")

type Primary struct {
	InlineLeXML
	XMLName xml.Name `xml:"primary"`
	Value   InnerXML
}

func (p *Primary) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var pp struct {
		InlineLeXML
		XMLName xml.Name `xml:"primary"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&pp, &start); err != nil {
		return errors.Wrap(ErrPrimaryUnknownElement, "failed to unmarshal data")
	}
	*p = Primary{
		XMLName: xml.Name{Local: "primary"},
		Value: InnerXML{
			Value: pp.Value,
		},
	}
	return nil
}
