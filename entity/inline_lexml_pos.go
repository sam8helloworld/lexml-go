package entity

import (
	"encoding/xml"
)

type SmallPos struct {
	InlineLeXML
	XMLName xml.Name `xml:"pos"`
	Value   PCDATA   `xml:",chardata"` // (#PCDATA)
}

type LargePos struct {
	InlineLeXML
	XMLName xml.Name `xml:"Pos"`
	Value   PCDATA   `xml:",chardata"` // (#PCDATA)
}

func (sp *SmallPos) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var smp struct {
		InlineLeXML
		XMLName xml.Name `xml:"pos"`
		Value   string   `xml:",chardata"` // (#PCDATA)
	}
	if err := d.DecodeElement(&smp, &start); err != nil {
		return err
	}
	*sp = SmallPos{
		XMLName: smp.XMLName,
		Value: PCDATA{
			Value: smp.Value,
		},
	}
	return nil
}

func (lp *LargePos) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var lrp struct {
		InlineLeXML
		XMLName xml.Name `xml:"POS"`
		Value   string   `xml:",chardata"` // (#PCDATA)
	}
	if err := d.DecodeElement(&lrp, &start); err != nil {
		return err
	}
	*lp = LargePos{
		XMLName: lrp.XMLName,
		Value: PCDATA{
			Value: lrp.Value,
		},
	}
	return nil
}
