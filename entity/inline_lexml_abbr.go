package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrAbbrUnknownElement = errors.New("abbr have unknown children")

type Abbr struct {
	InlineLeXML
	XMLName xml.Name `xml:"abbr"`
	Value   InnerXML
}

func (a *Abbr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var aa struct {
		InlineLeXML
		XMLName xml.Name `xml:"abbr"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&aa, &start); err != nil {
		return errors.Wrap(ErrAbbrUnknownElement, "failed to unmarshal data")
	}
	*a = Abbr{
		XMLName: xml.Name{Local: "abbr"},
		Value: InnerXML{
			Value: aa.Value,
		},
	}
	return nil
}
