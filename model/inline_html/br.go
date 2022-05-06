package model

import "encoding/xml"

type br struct {
	InlineHTML
	XMLName xml.Name `xml:"br"`
	Type    string   `xml:"type,attr"`
	Class   string   `xml:"class,attr"`
	Value   string   `xml:",chardata"`
}

type BR struct {
	InlineHTML
	XMLName xml.Name `xml:"BR"`
	Type    string   `xml:"type,attr"`
	Class   string   `xml:"class,attr"`
	Value   string   `xml:",chardata"`
}
