package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrUdUnknownElement = errors.New("ud have unknown children")

type Ud struct {
	InlineLeXML
	XMLName xml.Name `xml:"ud"`
	Type    string   `xml:"type,attr"`
	Class   string   `xml:"class,attr"`
	Value   InnerXML
}

func (u *Ud) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var uu struct {
		InlineLeXML
		XMLName xml.Name `xml:"ud"`
		Type    string   `xml:"type,attr"`
		Class   string   `xml:"class,attr"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&uu, &start); err != nil {
		return errors.Wrap(ErrUdUnknownElement, "failed to unmarshal data")
	}
	*u = Ud{
		XMLName: xml.Name{Local: "ud"},
		Type:    uu.Type,
		Class:   uu.Class,
		Value: InnerXML{
			Value: uu.Value,
		},
	}
	return nil
}
