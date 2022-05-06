package model

import "encoding/xml"

type em struct {
	InlineHTML
	XMLName xml.Name `xml:"em"`
	Type    string   `xml:"type,attr"`
	Class   string   `xml:"class,attr"`
	Value   string   `xml:",chardata"`
}

type EM struct {
	InlineHTML
	XMLName xml.Name `xml:"EM"`
	Type    string   `xml:"type,attr"`
	Class   string   `xml:"class,attr"`
	Value   string   `xml:",chardata"`
}
