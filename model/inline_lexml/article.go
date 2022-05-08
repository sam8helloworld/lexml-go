package inline_lexml

import "encoding/xml"

type Article struct {
	InlineLeXML
	XMLName xml.Name `xml:"article"`
	Value   string   `xml:",chardata"` // (#PCDATA | %inline.html; | %inline.lexml;)*
}
