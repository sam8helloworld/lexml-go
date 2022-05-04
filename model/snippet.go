package model

import "encoding/xml"

type Snippet struct {
	XMLName  xml.Name `xml:"snippet"`
	PID      string   `xml:"pid,attr"`
	Type     string   `xml:"type,attr"`
	Headword string   `xml:"headword,attr"`
	Value    string   `xml:",innerxml"`
}
