package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrEtymUnknownElement = errors.New("etym have unknown children")

type Etym struct {
	InlineLeXML
	XMLName xml.Name `xml:"etym"`
	Value   InnerXML
}

func (e *Etym) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var ee struct {
		InlineLeXML
		XMLName xml.Name `xml:"etym"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&ee, &start); err != nil {
		return errors.Wrap(ErrEtymUnknownElement, "failed to unmarshal data")
	}
	*e = Etym{
		XMLName: xml.Name{Local: "etym"},
		Value: InnerXML{
			Value: ee.Value,
		},
	}
	return nil
}
