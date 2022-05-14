package entity

import "encoding/xml"

type Accent struct {
	InlineLeXML
	XMLName xml.Name `xml:"accent"`
	Type    string   `xml:"type,attr"`
	Value   string   `xml:",chardata"` // (#PCDATA | %inline.html; | %inline.lexml;)*
}
