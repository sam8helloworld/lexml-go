package model

import "encoding/xml"

type B struct {
	InlineHTML
	XMLName xml.Name `xml:"b"`
	Type    string   `xml:"type,attr"`
	Class   string   `xml:"class,attr"`
	Value   string   `xml:",innerxml"`
}
