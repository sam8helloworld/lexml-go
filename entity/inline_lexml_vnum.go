package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrVnumUnknownElement = errors.New("vnum have unknown children")

type Vnum struct {
	InlineLeXML
	XMLName xml.Name `xml:"vnum"`
	Type    string   `xml:"type,attr"`
	Value   InnerXML
}

func (v *Vnum) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var vv struct {
		InlineLeXML
		XMLName xml.Name `xml:"vnum"`
		Type    string   `xml:"type,attr"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&vv, &start); err != nil {
		return errors.Wrap(ErrVnumUnknownElement, "failed to unmarshal data")
	}
	*v = Vnum{
		XMLName: xml.Name{Local: "vnum"},
		Type:    vv.Type,
		Value: InnerXML{
			Value: vv.Value,
		},
	}
	return nil
}
