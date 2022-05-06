package model

import "encoding/xml"

type Hidden struct {
	InlineHTML
	XMLName xml.Name `xml:"hidden"`
	Type    string   `xml:"type,attr"`
	Memo    string   `xml:"memo,attr"`
	Value   string   `xml:",chardata"`
}
