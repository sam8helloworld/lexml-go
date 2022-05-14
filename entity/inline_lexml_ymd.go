package entity

import "encoding/xml"

type Ymd struct {
	InlineLeXML
	XMLName xml.Name `xml:"ymd"`
	Type    string   `xml:"type,attr"`
	Value   string   `xml:",chardata"` // (#PCDATA | %inline.html; | %inline.lexml;)*
}
