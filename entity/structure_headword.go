package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrHeadwordUnknownElement = errors.New("headword have unknown children")

type Headword struct {
	Structure
	XMLName   xml.Name `xml:"headword"`
	Type      string   `xml:"type,attr"`
	Delimiter string   `xml:"delimiter,attr"`
	Value     InnerXML
}

func (h *Headword) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var hh struct {
		Structure
		XMLName   xml.Name `xml:"headword"`
		Type      string   `xml:"type,attr"`
		Delimiter string   `xml:"delimiter,attr"`
		Value     string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&hh, &start); err != nil {
		return errors.Wrap(ErrHeadwordUnknownElement, "failed to unmarshal data")
	}
	*h = Headword{
		XMLName:   xml.Name{Local: "headword"},
		Type:      hh.Type,
		Delimiter: hh.Delimiter,
		Value: InnerXML{
			Value: hh.Value,
		},
	}
	return nil
}
