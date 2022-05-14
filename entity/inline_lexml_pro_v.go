package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrProVUnknownElement = errors.New("pro-v have unknown children")

type ProV struct {
	InlineLeXML
	XMLName xml.Name `xml:"pro-v"`
	Value   InnerXML
}

func (p *ProV) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var pp struct {
		InlineLeXML
		XMLName xml.Name `xml:"pro-v"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&pp, &start); err != nil {
		return errors.Wrap(ErrProVUnknownElement, "failed to unmarshal data")
	}
	*p = ProV{
		XMLName: xml.Name{Local: "pro-v"},
		Value: InnerXML{
			Value: pp.Value,
		},
	}
	return nil
}
