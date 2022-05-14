package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrKlOkuriUnknownElement = errors.New("klokuri have unknown children")

type KlOkuri struct {
	InlineLeXML
	XMLName xml.Name `xml:"klokuri"`
	Value   InnerXML
}

func (k *KlOkuri) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var kk struct {
		InlineLeXML
		XMLName xml.Name `xml:"klokuri"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&kk, &start); err != nil {
		return errors.Wrap(ErrKlOkuriUnknownElement, "failed to unmarshal data")
	}
	*k = KlOkuri{
		XMLName: xml.Name{Local: "klokuri"},
		Value: InnerXML{
			Value: kk.Value,
		},
	}
	return nil
}
