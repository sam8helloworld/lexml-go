package structure

import (
	"encoding/xml"

	"github.com/pkg/errors"
	"github.com/sam8helloworld/lexml-go/model"
)

var ErrSubheadwordUnknownElement = errors.New("subheadword have unknown children")

type Subheadword struct {
	Structure
	XMLName   xml.Name `xml:"subheadword"`
	Type      string   `xml:"type,attr"`
	Delimiter string   `xml:"delimiter,attr"`
	Value     model.InnerXML
}

func (s *Subheadword) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var ss struct {
		Structure
		XMLName   xml.Name `xml:"subheadword"`
		Type      string   `xml:"type,attr"`
		Delimiter string   `xml:"delimiter,attr"`
		Value     string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&ss, &start); err != nil {
		return errors.Wrap(ErrSubheadwordUnknownElement, "failed to unmarshal data")
	}
	*s = Subheadword{
		XMLName:   xml.Name{Local: "subheadword"},
		Type:      ss.Type,
		Delimiter: ss.Delimiter,
		Value: model.InnerXML{
			Value: ss.Value,
		},
	}
	return nil
}
