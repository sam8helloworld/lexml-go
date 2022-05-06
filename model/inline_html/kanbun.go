package model

import "encoding/xml"

type Kanbun struct {
	InlineHTML
	XMLName xml.Name `xml:"kanbun"`
	Type    string   `xml:"type,attr"`
	SubID   string   `xml:"subid,attr"`
	Value   string   `xml:",chardata"`
}
