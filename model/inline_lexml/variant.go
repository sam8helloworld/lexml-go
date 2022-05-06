package model

import "encoding/xml"

type Variant struct {
	InlineLeXML
	XMLName xml.Name `xml:"variant"`
	Type    string   `xml:"type,attr"`
	Value   string   `xml:",chardata"`
}
