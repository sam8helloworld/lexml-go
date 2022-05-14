package entity

import "encoding/xml"

type ExMisc struct {
	InlineLeXML
	XMLName xml.Name `xml:"ex-misc"`
	Type    string   `xml:"type,attr"`
	Value   string   `xml:",chardata"` // (#PCDATA | %inline.html; | %inline.lexml;)*
}
