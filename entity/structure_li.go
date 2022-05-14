package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrLiUnknownElement = errors.New("li have unknown children")

type Li struct {
	Structure
	XMLName xml.Name `xml:"li"`
	SubID   string   `xml:"subid,attr"`
	Type    string   `xml:"type,attr"`
	Value   InnerXML
}

func (li *Li) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var l struct {
		Structure
		XMLName xml.Name `xml:"li"`
		SubID   string   `xml:"subid,attr"`
		Type    string   `xml:"type,attr"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&l, &start); err != nil {
		return errors.Wrap(ErrLiUnknownElement, "failed to unmarshal data")
	}
	*li = Li{
		XMLName: xml.Name{Local: "li"},
		SubID:   l.SubID,
		Type:    l.Type,
		Value: InnerXML{
			Value: l.Value,
		},
	}
	return nil
}
