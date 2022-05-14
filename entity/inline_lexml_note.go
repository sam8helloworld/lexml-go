package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrNoteUnknownElement = errors.New("note have unknown children")

type Note struct {
	InlineLeXML
	XMLName xml.Name `xml:"note"`
	Type    string   `xml:"type,attr"`
	Value   InnerXML
}

func (n *Note) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var nn struct {
		InlineLeXML
		XMLName xml.Name `xml:"note"`
		Type    string   `xml:"type,attr"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&nn, &start); err != nil {
		return errors.Wrap(ErrNoteUnknownElement, "failed to unmarshal data")
	}
	*n = Note{
		XMLName: xml.Name{Local: "note"},
		Value: InnerXML{
			Value: nn.Value,
		},
	}
	return nil
}
