package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrDataUnknownElement = errors.New("data have unknown children")

type Data struct {
	Structure
	XMLName xml.Name `xml:"data"`
	Type    string   `xml:"type,attr"`
	Value   InnerXML
}

func (data *Data) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var dd struct {
		Structure
		XMLName xml.Name `xml:"data"`
		Type    string   `xml:"type,attr"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&dd, &start); err != nil {
		return errors.Wrap(ErrDataUnknownElement, "failed to unmarshal data")
	}
	*data = Data{
		XMLName: xml.Name{Local: "data"},
		Type:    dd.Type,
		Value: InnerXML{
			Value: dd.Value,
		},
	}
	return nil
}
