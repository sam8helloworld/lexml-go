package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrVariantUnknownElement = errors.New("variant have unknown children")

type Variant struct {
	InlineLeXML
	XMLName xml.Name `xml:"variant"`
	Type    string   `xml:"type,attr"`
	Value   InnerXML
}

func (v *Variant) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var vv struct {
		InlineLeXML
		XMLName xml.Name `xml:"variant"`
		Type    string   `xml:"type,attr"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&vv, &start); err != nil {
		return errors.Wrap(ErrVariantUnknownElement, "failed to unmarshal data")
	}
	*v = Variant{
		XMLName: xml.Name{Local: "variant"},
		Type:    vv.Type,
		Value: InnerXML{
			Value: vv.Value,
		},
	}
	return nil
}
