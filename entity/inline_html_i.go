package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrIUnknownElement = errors.New("i have unknown children")

type SmallI struct {
	InlineHTML
	XMLName xml.Name `xml:"i"`
	Type    string   `xml:"type,attr"`
	Class   string   `xml:"class,attr"`
	Value   InnerXML
}

type LargeI struct {
	InlineHTML
	XMLName xml.Name `xml:"I"`
	Type    string   `xml:"type,attr"`
	Class   string   `xml:"class,attr"`
	Value   InnerXML
}

func (i *SmallI) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var ii struct {
		InlineHTML
		XMLName xml.Name `xml:"i"`
		Type    string   `xml:"type,attr"`
		Class   string   `xml:"class,attr"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&ii, &start); err != nil {
		return errors.Wrap(ErrIUnknownElement, "failed to unmarshal data")
	}
	*i = SmallI{
		XMLName: xml.Name{Local: "i"},
		Type:    ii.Type,
		Class:   ii.Class,
		Value: InnerXML{
			Value: ii.Value,
		},
	}
	return nil
}

func (i *LargeI) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var ii struct {
		InlineHTML
		XMLName xml.Name `xml:"I"`
		Type    string   `xml:"type,attr"`
		Class   string   `xml:"class,attr"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&ii, &start); err != nil {
		return errors.Wrap(ErrIUnknownElement, "failed to unmarshal data")
	}
	*i = LargeI{
		XMLName: xml.Name{Local: "I"},
		Type:    ii.Type,
		Class:   ii.Class,
		Value: InnerXML{
			Value: ii.Value,
		},
	}
	return nil
}
