package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrGLabelUnknownElement = errors.New("glabel have unknown children")

type GLabel struct {
	InlineLeXML
	XMLName xml.Name `xml:"glabel"`
	Type    string   `xml:"type,attr"`
	Value   InnerXML
}

func (g *GLabel) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var gg struct {
		InlineLeXML
		XMLName xml.Name `xml:"glabel"`
		Type    string   `xml:"type,attr"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&gg, &start); err != nil {
		return errors.Wrap(ErrGLabelUnknownElement, "failed to unmarshal data")
	}
	*g = GLabel{
		XMLName: xml.Name{Local: "glabel"},
		Type:    gg.Type,
		Value: InnerXML{
			Value: gg.Value,
		},
	}
	return nil
}
