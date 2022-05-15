package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrAUnknownElement = errors.New("a have unknown children")

type A struct {
	InlineHTML
	XMLName xml.Name `xml:"a"`
	Href    string   `xml:"href,attr"`
	Target  string   `xml:"target,attr"`
	Value   InnerXML
}

func (a *A) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var aa struct {
		InlineHTML
		XMLName xml.Name `xml:"a"`
		Href    string   `xml:"href,attr"`
		Target  string   `xml:"target,attr"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&aa, &start); err != nil {
		return errors.Wrap(ErrAUnknownElement, "failed to unmarshal data")
	}
	*a = A{
		XMLName: xml.Name{Local: "a"},
		Href:    aa.Href,
		Target:  aa.Target,
		Value: InnerXML{
			Value: aa.Value,
		},
	}
	return nil
}
