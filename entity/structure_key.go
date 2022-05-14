package entity

import (
	"encoding/xml"
)

type Key struct {
	Structure
	XMLName xml.Name `xml:"key"`
	Type    string   `xml:"type,attr"`
	Value   PCDATA   `xml:",chardata"` // (#PCDATA)
}

func (k *Key) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var key struct {
		Structure
		XMLName xml.Name `xml:"key"`
		Type    string   `xml:"type,attr"`
		Value   string   `xml:",chardata"` // (#PCDATA)
	}
	if err := d.DecodeElement(&key, &start); err != nil {
		return err
	}
	*k = Key{
		XMLName: key.XMLName,
		Type:    key.Type,
		Value: PCDATA{
			Value: key.Value,
		},
	}
	return nil
}
