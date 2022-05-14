package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrCnUnknownElement = errors.New("cn have unknown children")

type Cn struct {
	InlineLeXML
	XMLName xml.Name `xml:"cn"`
	Value   InnerXML
}

func (c *Cn) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var cc struct {
		InlineLeXML
		XMLName xml.Name `xml:"cn"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&cc, &start); err != nil {
		return errors.Wrap(ErrCnUnknownElement, "failed to unmarshal data")
	}
	*c = Cn{
		XMLName: xml.Name{Local: "cn"},
		Value: InnerXML{
			Value: cc.Value,
		},
	}
	return nil
}
