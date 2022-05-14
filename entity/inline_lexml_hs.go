package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrHsUnknownElement = errors.New("hs have unknown children")

type Hs struct {
	InlineLeXML
	XMLName xml.Name `xml:"hs"`
	Value   InnerXML
}

func (h *Hs) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var hh struct {
		InlineLeXML
		XMLName xml.Name `xml:"hs"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&hh, &start); err != nil {
		return errors.Wrap(ErrHsUnknownElement, "failed to unmarshal data")
	}
	*h = Hs{
		XMLName: xml.Name{Local: "hs"},
		Value: InnerXML{
			Value: hh.Value,
		},
	}
	return nil
}
