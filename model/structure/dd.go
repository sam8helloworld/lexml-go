package structure

import (
	"encoding/xml"

	"github.com/pkg/errors"
	"github.com/sam8helloworld/lexml-go/model"
)

var ErrDdUnknownElement = errors.New("dd have unknown children")

type Dd struct {
	Structure
	XMLName xml.Name `xml:"dd"`
	SubID   string   `xml:"subid,attr"`
	Type    string   `xml:"type,attr"`
	Value   model.InnerXML
}

func (dd *Dd) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var ddd struct {
		Structure
		XMLName xml.Name `xml:"dd"`
		SubID   string   `xml:"subid,attr"`
		Type    string   `xml:"type,attr"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&ddd, &start); err != nil {
		return errors.Wrap(ErrDdUnknownElement, "failed to unmarshal data")
	}
	*dd = Dd{
		XMLName: xml.Name{Local: "dd"},
		SubID:   ddd.SubID,
		Type:    ddd.Type,
		Value: model.InnerXML{
			Value: ddd.Value,
		},
	}
	return nil
}
