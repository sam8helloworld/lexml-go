package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrIpaUnknownElement = errors.New("ipa have unknown children")

type IPA struct {
	InlineLeXML
	XMLName xml.Name `xml:"ipa"`
	Value   InnerXML
}

func (i *IPA) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var ii struct {
		InlineLeXML
		XMLName xml.Name `xml:"ipa"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&ii, &start); err != nil {
		return errors.Wrap(ErrIpaUnknownElement, "failed to unmarshal data")
	}
	*i = IPA{
		XMLName: xml.Name{Local: "ipa"},
		Value: InnerXML{
			Value: ii.Value,
		},
	}
	return nil
}
