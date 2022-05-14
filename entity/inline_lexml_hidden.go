package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrHiddenUnknownElement = errors.New("hidden have unknown children")

type Hidden struct {
	InlineLeXML
	XMLName xml.Name `xml:"hidden"`
	Type    string   `xml:"type,attr"`
	Memo    string   `xml:"memo,attr"`
	Value   InnerXML
}

func (h *Hidden) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var hh struct {
		InlineLeXML
		XMLName xml.Name `xml:"hidden"`
		Type    string   `xml:"type,attr"`
		Memo    string   `xml:"memo,attr"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&hh, &start); err != nil {
		return errors.Wrap(ErrHiddenUnknownElement, "failed to unmarshal data")
	}
	*h = Hidden{
		XMLName: xml.Name{Local: "hidden"},
		Type:    hh.Type,
		Memo:    hh.Memo,
		Value: InnerXML{
			Value: hh.Value,
		},
	}
	return nil
}
