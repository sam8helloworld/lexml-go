package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrSLabelUnknownElement = errors.New("slabel have unknown children")

type SLabel struct {
	InlineLeXML
	XMLName xml.Name `xml:"slabel"`
	Type    string   `xml:"type,attr"`
	Genre   string   `xml:"genre,attr"`
	Code    string   `xml:"code,attr"`
	Value   InnerXML
}

func (s *SLabel) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var ss struct {
		InlineLeXML
		XMLName xml.Name `xml:"slabel"`
		Type    string   `xml:"type,attr"`
		Genre   string   `xml:"genre,attr"`
		Code    string   `xml:"code,attr"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&ss, &start); err != nil {
		return errors.Wrap(ErrSLabelUnknownElement, "failed to unmarshal data")
	}
	*s = SLabel{
		XMLName: xml.Name{Local: "slabel"},
		Type:    ss.Type,
		Genre:   ss.Genre,
		Code:    ss.Code,
		Value: InnerXML{
			Value: ss.Value,
		},
	}
	return nil
}
