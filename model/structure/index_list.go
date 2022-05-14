package structure

import (
	"encoding/xml"

	"github.com/pkg/errors"
	"github.com/sam8helloworld/lexml-go/model"
)

var ErrIndexListUnknownElement = errors.New("indexlist have unknown children")

type IndexList struct {
	Structure
	XMLName xml.Name `xml:"indexlist"`
	Type    string   `xml:"type,attr"`
	Value   model.InnerXML
}

func (i *IndexList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var ii struct {
		Structure
		XMLName xml.Name `xml:"indexlist"`
		Type    string   `xml:"type,attr"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&ii, &start); err != nil {
		return errors.Wrap(ErrIndexListUnknownElement, "failed to unmarshal data")
	}
	*i = IndexList{
		XMLName: xml.Name{Local: "indexlist"},
		Type:    ii.Type,
		Value: model.InnerXML{
			Value: ii.Value,
		},
	}
	return nil
}
