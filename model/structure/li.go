package structure

import (
	"encoding/xml"

	"github.com/pkg/errors"
	"github.com/sam8helloworld/lexml-go/model"
)

var ErrLiUnknownElement = errors.New("li have unknown children")

type Li struct {
	Structure
	XMLName xml.Name `xml:"li"`
	SubID   string   `xml:"subid,attr"`
	Type    string   `xml:"type,attr"`
	Value   model.InnerXML
}

func (li *Li) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var l struct {
		Structure
		XMLName xml.Name `xml:"li"`
		SubID   string   `xml:"subid,attr"`
		Type    string   `xml:"type,attr"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&l, &start); err != nil {
		return errors.Wrap(ErrLiUnknownElement, "failed to unmarshal data")
	}
	*li = Li{
		XMLName: xml.Name{Local: "li"},
		SubID:   l.SubID,
		Type:    l.Type,
		Value: model.InnerXML{
			Value: l.Value,
		},
	}
	return nil
}
