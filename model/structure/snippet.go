package model

import "encoding/xml"

type Snippet struct {
	Structure
	XMLName  xml.Name `xml:"snippet"`
	PID      string   `xml:"pid,attr"`
	Type     string   `xml:"type,attr"`
	Headword string   `xml:"headword,attr"`
	Value    string   `xml:",chardata"`
}
