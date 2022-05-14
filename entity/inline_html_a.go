package entity

import "encoding/xml"

type A struct {
	InlineHTML
	XMLName xml.Name `xml:"a"`
	Href    string   `xml:"href,attr"`
	Target  string   `xml:"target,attr"`
	Value   string   `xml:",chardata"` // (#PCDATA | %inline.html; | %inline.lexml;)*
}
