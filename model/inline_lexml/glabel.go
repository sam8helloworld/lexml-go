package inline_lexml

import "encoding/xml"

type GLabel struct {
	InlineLeXML
	XMLName xml.Name `xml:"glabel"`
	Type    string   `xml:"type,attr"`
	Value   string   `xml:",chardata"` // (#PCDATA | %inline.html; | %inline.lexml;)*
}
