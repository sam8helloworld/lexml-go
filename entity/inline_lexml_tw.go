package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrTwUnknownElement = errors.New("tw have unknown children")

type Tw struct {
	InlineLeXML
	XMLName xml.Name `xml:"tw"`
	Value   InnerXML
}

func (t *Tw) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var tt struct {
		InlineLeXML
		XMLName xml.Name `xml:"tw"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&tt, &start); err != nil {
		return errors.Wrap(ErrTwUnknownElement, "failed to unmarshal data")
	}
	*t = Tw{
		XMLName: xml.Name{Local: "tw"},

		Value: InnerXML{
			Value: tt.Value,
		},
	}
	return nil
}
