package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrProNUnknownElement = errors.New("pro-n have unknown children")

type ProN struct {
	InlineLeXML
	XMLName xml.Name `xml:"pro-n"`
	Value   InnerXML
}

func (p *ProN) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var pp struct {
		InlineLeXML
		XMLName xml.Name `xml:"pro-n"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&pp, &start); err != nil {
		return errors.Wrap(ErrProNUnknownElement, "failed to unmarshal data")
	}
	*p = ProN{
		XMLName: xml.Name{Local: "pro-n"},
		Value: InnerXML{
			Value: pp.Value,
		},
	}
	return nil
}
