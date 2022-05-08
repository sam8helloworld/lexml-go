package inline_lexml

import (
	"encoding/xml"

	"github.com/sam8helloworld/lexml-go/model/pcdata"
)

type Gix struct {
	InlineLeXML
	XMLName xml.Name      `xml:"gix"`
	Alt     string        `xml:"alt,attr"`
	Type    string        `xml:"type,attr"`
	AltImg  string        `xml:"altimg,attr"`
	Value   pcdata.PCDATA `xml:",chardata"` // (#PCDATA)
}

func (g *Gix) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var gix struct {
		InlineLeXML
		XMLName xml.Name `xml:"gix"`
		Alt     string   `xml:"alt,attr"`
		Type    string   `xml:"type,attr"`
		AltImg  string   `xml:"altimg,attr"`
		Value   string   `xml:",chardata"` // (#PCDATA)
	}
	if err := d.DecodeElement(&gix, &start); err != nil {
		return err
	}
	*g = Gix{
		XMLName: gix.XMLName,
		Alt:     gix.Alt,
		Type:    gix.Type,
		AltImg:  gix.AltImg,
		Value: pcdata.PCDATA{
			Value: gix.Value,
		},
	}
	return nil
}
