package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrKlRtUnknownElement = errors.New("klrt have unknown children")

type KlRt struct {
	InlineLeXML
	XMLName xml.Name `xml:"klrt"`
	Value   InnerXML
}

func (k *KlRt) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var kk struct {
		InlineLeXML
		XMLName xml.Name `xml:"klrt"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&kk, &start); err != nil {
		return errors.Wrap(ErrKlRtUnknownElement, "failed to unmarshal data")
	}
	*k = KlRt{
		XMLName: xml.Name{Local: "klrt"},
		Value: InnerXML{
			Value: kk.Value,
		},
	}
	return nil
}
