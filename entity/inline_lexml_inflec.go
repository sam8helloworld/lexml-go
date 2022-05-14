package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrInflecUnknownElement = errors.New("inflec have unknown children")

type Inflec struct {
	InlineLeXML
	XMLName xml.Name `xml:"inflec"`
	Value   InnerXML
}

func (i *Inflec) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var ii struct {
		InlineLeXML
		XMLName xml.Name `xml:"inflec"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&ii, &start); err != nil {
		return errors.Wrap(ErrInflecUnknownElement, "failed to unmarshal data")
	}
	*i = Inflec{
		XMLName: xml.Name{Local: "inflec"},
		Value: InnerXML{
			Value: ii.Value,
		},
	}
	return nil
}
