package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrKrUnknownElement = errors.New("kr have unknown children")

type Kr struct {
	InlineLeXML
	XMLName xml.Name `xml:"kr"`
	Value   InnerXML
}

func (k *Kr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var kk struct {
		InlineLeXML
		XMLName xml.Name `xml:"kr"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&kk, &start); err != nil {
		return errors.Wrap(ErrKrUnknownElement, "failed to unmarshal data")
	}
	*k = Kr{
		XMLName: xml.Name{Local: "kr"},
		Value: InnerXML{
			Value: kk.Value,
		},
	}
	return nil
}
