package model

import "encoding/xml"

type Ymd struct {
	InlineLeXML
	XMLName xml.Name `xml:"ymd"`
	Type    string   `xml:"type,attr"`
	Value   string   `xml:",chardata"`
}
