package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrSpelloutUnknownElement = errors.New("spellout have unknown children")

type Spellout struct {
	InlineLeXML
	XMLName xml.Name `xml:"spellout"`
	Type    string   `xml:"type,attr"`
	Org     string   `xml:"org,attr"`
	Value   InnerXML
}

func (s *Spellout) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var ss struct {
		InlineLeXML
		XMLName xml.Name `xml:"spellout"`
		Type    string   `xml:"type,attr"`
		Org     string   `xml:"org,attr"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&ss, &start); err != nil {
		return errors.Wrap(ErrSpelloutUnknownElement, "failed to unmarshal data")
	}
	*s = Spellout{
		XMLName: xml.Name{Local: "spellout"},
		Type:    ss.Type,
		Org:     ss.Org,
		Value: InnerXML{
			Value: ss.Value,
		},
	}
	return nil
}
