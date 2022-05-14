package entity

import "encoding/xml"

type Ggk struct {
	InlineLeXML
	XMLName xml.Name `xml:"glabel"`
	Class   string   `xml:"class,attr"`
	Yomi    string   `xml:"yomi,attr"`
	Value   string   `xml:",chardata"` // (#PCDATA | %inline.html; | %inline.lexml;)*
}
