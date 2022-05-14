package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrFullFormUnknownElement = errors.New("fullform have unknown children")

type FullForm struct {
	InlineLeXML
	XMLName xml.Name `xml:"fullform"`
	Value   InnerXML
}

func (f *FullForm) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var ff struct {
		InlineLeXML
		XMLName xml.Name `xml:"fullform"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&ff, &start); err != nil {
		return errors.Wrap(ErrFullFormUnknownElement, "failed to unmarshal data")
	}
	*f = FullForm{
		XMLName: xml.Name{Local: "fullform"},
		Value: InnerXML{
			Value: ff.Value,
		},
	}
	return nil
}
