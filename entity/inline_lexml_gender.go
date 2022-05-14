package entity

import (
	"encoding/xml"
)

type SmallGender struct {
	InlineLeXML
	XMLName xml.Name `xml:"gender"`
	Value   PCDATA   `xml:",chardata"` // (#PCDATA)
}

type LargeGender struct {
	InlineLeXML
	XMLName xml.Name `xml:"GENDER"`
	Value   PCDATA   `xml:",chardata"` // (#PCDATA)
}

func (sg *SmallGender) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var smg struct {
		InlineLeXML
		XMLName xml.Name `xml:"gender"`
		Value   string   `xml:",chardata"` // (#PCDATA)
	}
	if err := d.DecodeElement(&smg, &start); err != nil {
		return err
	}
	*sg = SmallGender{
		XMLName: smg.XMLName,
		Value: PCDATA{
			Value: smg.Value,
		},
	}
	return nil
}

func (lg *LargeGender) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var lag struct {
		InlineLeXML
		XMLName xml.Name `xml:"GENDER"`
		Value   string   `xml:",chardata"` // (#PCDATA)
	}
	if err := d.DecodeElement(&lag, &start); err != nil {
		return err
	}
	*lg = LargeGender{
		XMLName: lag.XMLName,
		Value: PCDATA{
			Value: lag.Value,
		},
	}
	return nil
}
