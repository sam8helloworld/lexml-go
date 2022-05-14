package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrExTransUnknownElement = errors.New("ex-trans have unknown children")

type ExTrans struct {
	InlineLeXML
	XMLName xml.Name `xml:"ex-trans"`
	Lang    string   `xml:"lang,attr"`
	Type    string   `xml:"type,attr"`
	Value   InnerXML
}

func (e *ExTrans) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var ee struct {
		InlineLeXML
		XMLName xml.Name `xml:"ex-trans"`
		Lang    string   `xml:"lang,attr"`
		Type    string   `xml:"type,attr"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&ee, &start); err != nil {
		return errors.Wrap(ErrExTransUnknownElement, "failed to unmarshal data")
	}
	*e = ExTrans{
		XMLName: xml.Name{Local: "ex-trans"},
		Lang:    ee.Lang,
		Type:    ee.Type,
		Value: InnerXML{
			Value: ee.Value,
		},
	}
	return nil
}
