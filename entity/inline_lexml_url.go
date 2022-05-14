package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrUrlUnknownElement = errors.New("url have unknown children")

type Url struct {
	InlineLeXML
	XMLName xml.Name `xml:"url"`
	Type    string   `xml:"type,attr"`
	Value   InnerXML
}

func (u *Url) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var uu struct {
		InlineLeXML
		XMLName xml.Name `xml:"url"`
		Type    string   `xml:"type,attr"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&uu, &start); err != nil {
		return errors.Wrap(ErrUrlUnknownElement, "failed to unmarshal data")
	}
	*u = Url{
		XMLName: xml.Name{Local: "url"},
		Type:    uu.Type,
		Value: InnerXML{
			Value: uu.Value,
		},
	}
	return nil
}
