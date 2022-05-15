package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrSubUnknownElement = errors.New("sub have unknown children")

type SmallSub struct {
	InlineHTML
	XMLName xml.Name `xml:"sub"`
	Value   InnerXML
}

type LargeSub struct {
	InlineHTML
	XMLName xml.Name `xml:"SUB"`
	Value   InnerXML
}

func (s *SmallSub) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var ss struct {
		InlineHTML
		XMLName xml.Name `xml:"sub"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&ss, &start); err != nil {
		return errors.Wrap(ErrSubUnknownElement, "failed to unmarshal data")
	}
	*s = SmallSub{
		XMLName: xml.Name{Local: "sub"},
		Value: InnerXML{
			Value: ss.Value,
		},
	}
	return nil
}

func (s *LargeSub) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var ss struct {
		InlineHTML
		XMLName xml.Name `xml:"SUB"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&ss, &start); err != nil {
		return errors.Wrap(ErrSubUnknownElement, "failed to unmarshal data")
	}
	*s = LargeSub{
		XMLName: xml.Name{Local: "SUB"},
		Value: InnerXML{
			Value: ss.Value,
		},
	}
	return nil
}
