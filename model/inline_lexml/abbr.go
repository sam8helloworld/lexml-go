package inline_lexml

import "encoding/xml"

type Abbr struct {
	InlineLeXML
	XMLName xml.Name `xml:"abbr"`
	Value   string   `xml:",chardata"` // (#PCDATA | %inline.html; | %inline.lexml;)*
}
