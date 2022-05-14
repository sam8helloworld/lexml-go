package entity

import (
	"encoding/xml"
)

type Snippet struct {
	Structure
	XMLName  xml.Name `xml:"snippet"`
	PID      string   `xml:"pid,attr"`
	Type     string   `xml:"type,attr"`
	Headword string   `xml:"headword,attr"`
	Value    PCDATA
}

func (s *Snippet) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var sn struct {
		Structure
		XMLName  xml.Name `xml:"snippet"`
		PID      string   `xml:"pid,attr"`
		Type     string   `xml:"type,attr"`
		Headword string   `xml:"headword,attr"`
		Value    string   `xml:",innerxml"` //  (#PCDATA)
	}
	if err := d.DecodeElement(&sn, &start); err != nil {
		return err
	}
	*s = Snippet{
		XMLName:  sn.XMLName,
		PID:      sn.PID,
		Type:     sn.Type,
		Headword: sn.Headword,
		Value: PCDATA{
			Value: sn.Value,
		},
	}
	return nil
}
