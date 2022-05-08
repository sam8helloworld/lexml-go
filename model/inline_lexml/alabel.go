package inline_lexml

import "encoding/xml"

type ALabel struct {
	InlineLeXML
	XMLName xml.Name `xml:"alabel"`
	Type    string   `xml:"type,attr"`
	Code    string   `xml:"code,attr"`
	Value   string   `xml:",chardata"` // (#PCDATA | %inline.html; | %inline.lexml;)*
}
