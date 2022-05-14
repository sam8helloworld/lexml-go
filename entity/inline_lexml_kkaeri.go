package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrKKaeriUnknownElement = errors.New("kkaeri have unknown children")

type KKaeri struct {
	InlineLeXML
	XMLName xml.Name `xml:"kkaeri"`
	Value   InnerXML
}

func (k *KKaeri) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var kk struct {
		InlineLeXML
		XMLName xml.Name `xml:"kkaeri"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&kk, &start); err != nil {
		return errors.Wrap(ErrKKaeriUnknownElement, "failed to unmarshal data")
	}
	*k = KKaeri{
		XMLName: xml.Name{Local: "kkaeri"},
		Value: InnerXML{
			Value: kk.Value,
		},
	}
	return nil
}
