package inline_lexml

import (
	"encoding/xml"

	"github.com/sam8helloworld/lexml-go/model/pcdata"
)

type SmallGender struct {
	InlineLeXML
	XMLName xml.Name      `xml:"gender"`
	Value   pcdata.PCDATA `xml:",chardata"` // (#PCDATA)
}

type LargeGender struct {
	InlineLeXML
	XMLName xml.Name      `xml:"GENDER"`
	Value   pcdata.PCDATA `xml:",chardata"` // (#PCDATA)
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
		Value: pcdata.PCDATA{
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
		Value: pcdata.PCDATA{
			Value: lag.Value,
		},
	}
	return nil
}
