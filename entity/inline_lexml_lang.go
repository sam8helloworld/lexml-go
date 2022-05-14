package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrLangUnknownElement = errors.New("kigo have unknown children")

type Lang struct {
	InlineLeXML
	XMLName xml.Name `xml:"lang"`
	Type    string   `xml:"type,attr"`
	Value   InnerXML
}

func (l *Lang) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var ll struct {
		InlineLeXML
		XMLName xml.Name `xml:"lang"`
		Type    string   `xml:"type,attr"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&ll, &start); err != nil {
		return errors.Wrap(ErrLangUnknownElement, "failed to unmarshal data")
	}
	*l = Lang{
		XMLName: xml.Name{Local: "lang"},
		Type:    ll.Type,
		Value: InnerXML{
			Value: ll.Value,
		},
	}
	return nil
}
