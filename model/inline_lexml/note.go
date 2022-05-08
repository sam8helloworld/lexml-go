package inline_lexml

import "encoding/xml"

type Note struct {
	InlineLeXML
	XMLName xml.Name `xml:"note"`
	Type    string   `xml:"type,attr"`
	Value   string   `xml:",chardata"` // (#PCDATA | %inline.html; | %inline.lexml;)*
}
