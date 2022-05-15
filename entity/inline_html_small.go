package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrSmallUnknownElement = errors.New("small have unknown children")

type Small struct {
	InlineHTML
	XMLName xml.Name `xml:"small"`
	Value   InnerXML
}

func (s *Small) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var ss struct {
		InlineHTML
		XMLName xml.Name `xml:"small"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&ss, &start); err != nil {
		return errors.Wrap(ErrSmallUnknownElement, "failed to unmarshal data")
	}
	*s = Small{
		XMLName: xml.Name{Local: "small"},
		Value: InnerXML{
			Value: ss.Value,
		},
	}
	return nil
}
