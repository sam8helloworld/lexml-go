package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrSpanUnknownElement = errors.New("span have unknown children")

type Span struct {
	InlineHTML
	XMLName xml.Name `xml:"span"`
	Type    string   `xml:"type,attr"`
	Class   string   `xml:"class,attr"`
	Ph      string   `xml:"ph,attr"`
	Value   InnerXML
}

func (s *Span) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var ss struct {
		InlineHTML
		XMLName xml.Name `xml:"span"`
		Type    string   `xml:"type,attr"`
		Class   string   `xml:"class,attr"`
		Ph      string   `xml:"ph,attr"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&ss, &start); err != nil {
		return errors.Wrap(ErrSpanUnknownElement, "failed to unmarshal data")
	}
	*s = Span{
		XMLName: xml.Name{Local: "span"},
		Type:    ss.Type,
		Class:   ss.Class,
		Ph:      ss.Ph,
		Value: InnerXML{
			Value: ss.Value,
		},
	}
	return nil
}
