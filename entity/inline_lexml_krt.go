package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrKRtUnknownElement = errors.New("krt have unknown children")

type KRt struct {
	InlineLeXML
	XMLName xml.Name `xml:"krt"`
	Value   InnerXML
}

func (k *KRt) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var kk struct {
		InlineLeXML
		XMLName xml.Name `xml:"krt"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&kk, &start); err != nil {
		return errors.Wrap(ErrKRtUnknownElement, "failed to unmarshal data")
	}
	*k = KRt{
		XMLName: xml.Name{Local: "krt"},
		Value: InnerXML{
			Value: kk.Value,
		},
	}
	return nil
}
