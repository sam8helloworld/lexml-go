package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrSupUnknownElement = errors.New("sup have unknown children")

type SmallSup struct {
	InlineHTML
	XMLName xml.Name `xml:"sup"`
	Value   InnerXML
}

type LargeSup struct {
	InlineHTML
	XMLName xml.Name `xml:"SUP"`
	Value   InnerXML
}

func (s *SmallSup) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var ss struct {
		InlineHTML
		XMLName xml.Name `xml:"sup"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&ss, &start); err != nil {
		return errors.Wrap(ErrSupUnknownElement, "failed to unmarshal data")
	}
	*s = SmallSup{
		XMLName: xml.Name{Local: "sup"},
		Value: InnerXML{
			Value: ss.Value,
		},
	}
	return nil
}

func (s *LargeSup) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var ss struct {
		InlineHTML
		XMLName xml.Name `xml:"SUP"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&ss, &start); err != nil {
		return errors.Wrap(ErrSupUnknownElement, "failed to unmarshal data")
	}
	*s = LargeSup{
		XMLName: xml.Name{Local: "SUP"},
		Value: InnerXML{
			Value: ss.Value,
		},
	}
	return nil
}
