package inline_lexml

import "encoding/xml"

type Sc struct {
	InlineLeXML
	XMLName xml.Name `xml:"sc"`
	Value   string   `xml:",chardata"` // (#PCDATA | %inline.html; | %inline.lexml;)*
}
