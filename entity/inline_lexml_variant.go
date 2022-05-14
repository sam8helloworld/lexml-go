package entity

import "encoding/xml"

type Variant struct {
	InlineLeXML
	XMLName xml.Name `xml:"variant"`
	Type    string   `xml:"type,attr"`
	Value   string   `xml:",chardata"` // (#PCDATA | %inline.html; | %inline.lexml;)*
}
