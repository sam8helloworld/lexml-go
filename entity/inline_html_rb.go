package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrRbUnknownElement = errors.New("rb have unknown children")

type Rb struct {
	InlineHTML
	XMLName xml.Name `xml:"rb"`
	Value   InnerXML
}

func (r *Rb) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var rr struct {
		InlineHTML
		XMLName xml.Name `xml:"rb"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&rr, &start); err != nil {
		return errors.Wrap(ErrRbUnknownElement, "failed to unmarshal data")
	}
	*r = Rb{
		XMLName: xml.Name{Local: "rb"},
		Value: InnerXML{
			Value: rr.Value,
		},
	}
	return nil
}
