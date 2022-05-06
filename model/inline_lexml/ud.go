package model

import "encoding/xml"

type Ud struct {
	InlineLeXML
	XMLName xml.Name `xml:"ud"`
	Type    string   `xml:"type,attr"`
	Class   string   `xml:"class,attr"`
	Value   string   `xml:",chardata"`
}
