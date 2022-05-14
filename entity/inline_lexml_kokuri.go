package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrKOkuriUnknownElement = errors.New("kokuri have unknown children")

type KOkuri struct {
	InlineLeXML
	XMLName xml.Name `xml:"kokuri"`
	Value   InnerXML
}

func (k *KOkuri) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var kk struct {
		InlineLeXML
		XMLName xml.Name `xml:"kokuri"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&kk, &start); err != nil {
		return errors.Wrap(ErrKOkuriUnknownElement, "failed to unmarshal data")
	}
	*k = KOkuri{
		XMLName: xml.Name{Local: "kokuri"},
		Value: InnerXML{
			Value: kk.Value,
		},
	}
	return nil
}
