package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrOblqUnknownElement = errors.New("oblq have unknown children")

type Oblq struct {
	InlineLeXML
	XMLName xml.Name `xml:"oblq"`
	Value   InnerXML
}

func (o *Oblq) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var oo struct {
		InlineLeXML
		XMLName xml.Name `xml:"oblq"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&oo, &start); err != nil {
		return errors.Wrap(ErrOblqUnknownElement, "failed to unmarshal data")
	}
	*o = Oblq{
		XMLName: xml.Name{Local: "oblq"},
		Value: InnerXML{
			Value: oo.Value,
		},
	}
	return nil
}
