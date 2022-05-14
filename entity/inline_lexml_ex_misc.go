package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrExMiscUnknownElement = errors.New("ex-misc have unknown children")

type ExMisc struct {
	InlineLeXML
	XMLName xml.Name `xml:"ex-misc"`
	Type    string   `xml:"type,attr"`
	Value   InnerXML
}

func (e *ExMisc) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var ee struct {
		InlineLeXML
		XMLName xml.Name `xml:"ex-misc"`
		Type    string   `xml:"type,attr"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&ee, &start); err != nil {
		return errors.Wrap(ErrExMiscUnknownElement, "failed to unmarshal data")
	}
	*e = ExMisc{
		XMLName: xml.Name{Local: "ex-misc"},
		Type:    ee.Type,
		Value: InnerXML{
			Value: ee.Value,
		},
	}
	return nil
}
