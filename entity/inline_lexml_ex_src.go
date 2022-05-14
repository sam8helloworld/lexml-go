package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrExSrcUnknownElement = errors.New("ex-src have unknown children")

type ExSrc struct {
	InlineLeXML
	XMLName xml.Name `xml:"ex-src"`
	RefID   string   `xml:"refid,attr"`
	Type    string   `xml:"type,attr"`
	Value   InnerXML
}

func (e *ExSrc) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var ee struct {
		InlineLeXML
		XMLName xml.Name `xml:"ex-src"`
		RefID   string   `xml:"refid,attr"`
		Type    string   `xml:"type,attr"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&ee, &start); err != nil {
		return errors.Wrap(ErrExSrcUnknownElement, "failed to unmarshal data")
	}
	*e = ExSrc{
		XMLName: xml.Name{Local: "ex-src"},
		RefID:   ee.RefID,
		Type:    ee.Type,
		Value: InnerXML{
			Value: ee.Value,
		},
	}
	return nil
}
