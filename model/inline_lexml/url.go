package inline_lexml

import "encoding/xml"

type Url struct {
	InlineLeXML
	XMLName xml.Name `xml:"url"`
	Type    string   `xml:"type,attr"`
	Value   string   `xml:",chardata"` // (#PCDATA | %inline.html; | %inline.lexml;)*
}
