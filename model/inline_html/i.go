package model

import "encoding/xml"

type i struct {
	InlineHTML
	XMLName xml.Name `xml:"i"`
	Type    string   `xml:"type,attr"`
	Class   string   `xml:"class,attr"`
	Value   string   `xml:",chardata"`
}

type I struct {
	InlineHTML
	XMLName xml.Name `xml:"I"`
	Type    string   `xml:"type,attr"`
	Class   string   `xml:"class,attr"`
	Value   string   `xml:",chardata"`
}
