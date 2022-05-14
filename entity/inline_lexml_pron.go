package entity

import (
	"encoding/xml"
)

type Pron struct {
	InlineLeXML
	XMLName xml.Name `xml:"pron"`
	Value   PCDATA   `xml:",chardata"` // (#PCDATA)
}

func (p *Pron) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var pr struct {
		InlineLeXML
		XMLName xml.Name `xml:"pron"`
		Value   string   `xml:",chardata"` // (#PCDATA)
	}
	if err := d.DecodeElement(&pr, &start); err != nil {
		return err
	}
	*p = Pron{
		XMLName: pr.XMLName,
		Value: PCDATA{
			Value: pr.Value,
		},
	}
	return nil
}
