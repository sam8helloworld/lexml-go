package entity

import "encoding/xml"

type SmallB struct {
	InlineHTML
	XMLName xml.Name `xml:"b"`
	Type    string   `xml:"type,attr"`
	Class   string   `xml:"class,attr"`
	Value   string   `xml:",chardata"` // (#PCDATA | %inline.html; | %inline.lexml;)*
}

type LargeB struct {
	InlineHTML
	XMLName xml.Name `xml:"B"`
	Type    string   `xml:"type,attr"`
	Class   string   `xml:"class,attr"`
	Value   string   `xml:",chardata"` // (#PCDATA | %inline.html; | %inline.lexml;)*
}
