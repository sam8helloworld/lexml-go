package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrGgkUnknownElement = errors.New("ggk have unknown children")

type Ggk struct {
	InlineLeXML
	XMLName xml.Name `xml:"ggk"`
	Class   string   `xml:"class,attr"`
	Yomi    string   `xml:"yomi,attr"`
	Value   InnerXML
}

func (g *Ggk) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var gg struct {
		InlineLeXML
		XMLName xml.Name `xml:"ggk"`
		Class   string   `xml:"class,attr"`
		Yomi    string   `xml:"yomi,attr"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&gg, &start); err != nil {
		return errors.Wrap(ErrGgkUnknownElement, "failed to unmarshal data")
	}
	*g = Ggk{
		XMLName: xml.Name{Local: "ggk"},
		Class:   gg.Class,
		Yomi:    gg.Yomi,
		Value: InnerXML{
			Value: gg.Value,
		},
	}
	return nil
}
