package model

import "encoding/xml"

type Url struct {
	InlineLeXML
	XMLName xml.Name `xml:"url"`
	Type    string   `xml:"type,attr"`
	Value   string   `xml:",chardata"`
}
