package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrThUnknownElement = errors.New("th have unknown children")

type Th struct {
	Structure
	XMLName xml.Name `xml:"th"`
	ColSpan string   `xml:"colspan,attr"`
	RowSpan string   `xml:"rowspan,attr"`
	Value   InnerXML
}

func (t *Th) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var tt struct {
		Structure
		XMLName xml.Name `xml:"th"`
		ColSpan string   `xml:"colspan,attr"`
		RowSpan string   `xml:"rowspan,attr"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&tt, &start); err != nil {
		return errors.Wrap(ErrThUnknownElement, "failed to unmarshal data")
	}
	*t = Th{
		XMLName: xml.Name{Local: "th"},
		ColSpan: tt.ColSpan,
		RowSpan: tt.RowSpan,
		Value: InnerXML{
			Value: tt.Value,
		},
	}
	return nil
}
