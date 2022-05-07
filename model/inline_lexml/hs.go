package inline_lexml

import "encoding/xml"

type Hs struct {
	InlineLeXML
	XMLName xml.Name `xml:"hs"`
	Value   string   `xml:",chardata"` // (#PCDATA | %inline.html; | %inline.lexml;)*
}
