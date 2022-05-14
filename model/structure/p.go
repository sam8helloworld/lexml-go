package structure

import (
	"encoding/xml"

	"github.com/pkg/errors"
	"github.com/sam8helloworld/lexml-go/model"
)

var ErrPUnknownElement = errors.New("p have unknown children")

type P struct {
	Structure
	XMLName xml.Name `xml:"p"`
	SubID   string   `xml:"subid,attr"`
	Type    string   `xml:"type,attr"`
	Value   model.InnerXML
}

func (p *P) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var pp struct {
		Structure
		XMLName xml.Name `xml:"p"`
		SubID   string   `xml:"subid,attr"`
		Type    string   `xml:"type,attr"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&pp, &start); err != nil {
		return errors.Wrap(ErrPUnknownElement, "failed to unmarshal data")
	}
	*p = P{
		XMLName: xml.Name{Local: "p"},
		SubID:   pp.SubID,
		Type:    pp.Type,
		Value: model.InnerXML{
			Value: pp.Value,
		},
	}
	return nil
}
