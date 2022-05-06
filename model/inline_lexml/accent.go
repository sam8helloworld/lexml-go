package model

import "encoding/xml"

type Accent struct {
	InlineLeXML
	XMLName xml.Name `xml:"accent"`
	Type    string   `xml:"type,attr"`
	Value   string   `xml:",chardata"`
}
