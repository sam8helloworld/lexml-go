package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrExBodyUnknownElement = errors.New("ex-body have unknown children")

type ExBody struct {
	InlineLeXML
	XMLName xml.Name `xml:"ex-body"`
	Lang    string   `xml:"lang,attr"`
	Type    string   `xml:"type,attr"`
	Value   InnerXML
}

func (e *ExBody) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var ee struct {
		InlineLeXML
		XMLName xml.Name `xml:"ex-body"`
		Lang    string   `xml:"lang,attr"`
		Type    string   `xml:"type,attr"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&ee, &start); err != nil {
		return errors.Wrap(ErrExBodyUnknownElement, "failed to unmarshal data")
	}
	*e = ExBody{
		XMLName: xml.Name{Local: "ex-body"},
		Lang:    ee.Lang,
		Type:    ee.Type,
		Value: InnerXML{
			Value: ee.Value,
		},
	}
	return nil
}
