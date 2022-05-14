package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrYmdUnknownElement = errors.New("ymd have unknown children")

type Ymd struct {
	InlineLeXML
	XMLName xml.Name `xml:"ymd"`
	Type    string   `xml:"type,attr"`
	Value   InnerXML
}

func (y *Ymd) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var yy struct {
		InlineLeXML
		XMLName xml.Name `xml:"ymd"`
		Type    string   `xml:"type,attr"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&yy, &start); err != nil {
		return errors.Wrap(ErrYmdUnknownElement, "failed to unmarshal data")
	}
	*y = Ymd{
		XMLName: xml.Name{Local: "ymd"},
		Type:    yy.Type,
		Value: InnerXML{
			Value: yy.Value,
		},
	}
	return nil
}
