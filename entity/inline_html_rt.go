package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrRtUnknownElement = errors.New("rt have unknown children")

type Rt struct {
	InlineHTML
	XMLName xml.Name `xml:"rt"`
	Value   InnerXML
}

func (r *Rt) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var rr struct {
		InlineHTML
		XMLName xml.Name `xml:"rt"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&rr, &start); err != nil {
		return errors.Wrap(ErrRtUnknownElement, "failed to unmarshal data")
	}
	*r = Rt{
		XMLName: xml.Name{Local: "rt"},
		Value: InnerXML{
			Value: rr.Value,
		},
	}
	return nil
}
