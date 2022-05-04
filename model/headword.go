package model

import "encoding/xml"

type Headword struct {
	XMLName   xml.Name `xml:"headword"`
	Type      string   `xml:"type,attr"`
	Delimiter string   `xml:"delimiter,attr"`
	Value     string   `xml:",innerxml"`
}
