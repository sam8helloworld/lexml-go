package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrScUnknownElement = errors.New("sc have unknown children")

type Sc struct {
	InlineLeXML
	XMLName xml.Name `xml:"sc"`
	Value   InnerXML
}

func (s *Sc) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var ss struct {
		InlineLeXML
		XMLName xml.Name `xml:"sc"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&ss, &start); err != nil {
		return errors.Wrap(ErrScUnknownElement, "failed to unmarshal data")
	}
	*s = Sc{
		XMLName: xml.Name{Local: "sc"},
		Value: InnerXML{
			Value: ss.Value,
		},
	}
	return nil
}
