package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrKigoUnknownElement = errors.New("kigo have unknown children")

type Kigo struct {
	InlineLeXML
	XMLName xml.Name `xml:"kigo"`
	Value   InnerXML
}

func (k *Kigo) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var kk struct {
		InlineLeXML
		XMLName xml.Name `xml:"kigo"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&kk, &start); err != nil {
		return errors.Wrap(ErrKigoUnknownElement, "failed to unmarshal data")
	}
	*k = Kigo{
		XMLName: xml.Name{Local: "kigo"},
		Value: InnerXML{
			Value: kk.Value,
		},
	}
	return nil
}
