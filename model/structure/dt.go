package structure

import (
	"encoding/xml"

	"github.com/pkg/errors"
	"github.com/sam8helloworld/lexml-go/model"
)

var ErrDtUnknownElement = errors.New("dt have unknown children")

type Dt struct {
	Structure
	XMLName xml.Name `xml:"dt"`
	SubID   string   `xml:"subid,attr"`
	Type    string   `xml:"type,attr"`
	Value   model.InnerXML
}

func (dt *Dt) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var dtt struct {
		Structure
		XMLName xml.Name `xml:"dt"`
		SubID   string   `xml:"subid,attr"`
		Type    string   `xml:"type,attr"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&dtt, &start); err != nil {
		return errors.Wrap(ErrDtUnknownElement, "failed to unmarshal data")
	}
	*dt = Dt{
		XMLName: xml.Name{Local: "dt"},
		SubID:   dtt.SubID,
		Type:    dtt.Type,
		Value: model.InnerXML{
			Value: dtt.Value,
		},
	}
	return nil
}
