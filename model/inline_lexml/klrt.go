package inline_lexml

import "encoding/xml"

type KlRt struct {
	InlineLeXML
	XMLName xml.Name `xml:"klrt"`
	Value   string   `xml:",chardata"` // (#PCDATA | %inline.html; | %inline.lexml;)*
}
