package model

import "encoding/xml"

type Subheadword struct {
	XMLName   xml.Name `xml:"subheadword"`
	Type      string   `xml:"type,attr"`
	Delimiter string   `xml:"delimiter,attr"`
	Value     string   `xml:",chardata"`
}
