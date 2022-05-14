package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrJpUnknownElement = errors.New("jp have unknown children")

type Jp struct {
	InlineLeXML
	XMLName xml.Name `xml:"jp"`
	Value   InnerXML
}

func (j *Jp) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var jj struct {
		InlineLeXML
		XMLName xml.Name `xml:"jp"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&jj, &start); err != nil {
		return errors.Wrap(ErrJpUnknownElement, "failed to unmarshal data")
	}
	*j = Jp{
		XMLName: xml.Name{Local: "jp"},
		Value: InnerXML{
			Value: jj.Value,
		},
	}
	return nil
}
