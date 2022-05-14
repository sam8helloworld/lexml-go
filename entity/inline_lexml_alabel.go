package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrAlabelUnknownElement = errors.New("alabel have unknown children")

type ALabel struct {
	InlineLeXML
	XMLName xml.Name `xml:"alabel"`
	Type    string   `xml:"type,attr"`
	Code    string   `xml:"code,attr"`
	Value   InnerXML
}

func (a *ALabel) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var aa struct {
		InlineLeXML
		XMLName xml.Name `xml:"alabel"`
		Type    string   `xml:"type,attr"`
		Code    string   `xml:"code,attr"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&aa, &start); err != nil {
		return errors.Wrap(ErrAlabelUnknownElement, "failed to unmarshal data")
	}
	*a = ALabel{
		XMLName: xml.Name{Local: "alabel"},
		Type:    aa.Type,
		Code:    aa.Code,
		Value: InnerXML{
			Value: aa.Value,
		},
	}
	return nil
}
