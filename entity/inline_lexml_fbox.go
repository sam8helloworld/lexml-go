package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrFBoxUnknownElement = errors.New("fbox have unknown children")

type SmallFBox struct {
	InlineLeXML
	XMLName xml.Name `xml:"fbox"`
	Value   InnerXML
}

type LargeFBox struct {
	InlineLeXML
	XMLName xml.Name `xml:"FBOX"`
	Value   InnerXML
}

func (sf *SmallFBox) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var sff struct {
		InlineLeXML
		XMLName xml.Name `xml:"fbox"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&sff, &start); err != nil {
		return errors.Wrap(ErrFBoxUnknownElement, "failed to unmarshal data")
	}
	*sf = SmallFBox{
		XMLName: xml.Name{Local: "fbox"},
		Value: InnerXML{
			Value: sff.Value,
		},
	}
	return nil
}

func (lf *LargeFBox) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var lff struct {
		InlineLeXML
		XMLName xml.Name `xml:"FBOX"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&lff, &start); err != nil {
		return errors.Wrap(ErrFBoxUnknownElement, "failed to unmarshal data")
	}
	*lf = LargeFBox{
		XMLName: xml.Name{Local: "FBOX"},
		Value: InnerXML{
			Value: lff.Value,
		},
	}
	return nil
}
