package model

import "encoding/xml"

type u struct {
	InlineHTML
	XMLName xml.Name `xml:"u"`
	Type    string   `xml:"type,attr"`
	Class   string   `xml:"class,attr"`
	Value   string   `xml:",chardata"`
}

type U struct {
	InlineHTML
	XMLName xml.Name `xml:"U"`
	Type    string   `xml:"type,attr"`
	Class   string   `xml:"class,attr"`
	Value   string   `xml:",chardata"`
}
