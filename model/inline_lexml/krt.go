package inline_lexml

import "encoding/xml"

type KRt struct {
	InlineLeXML
	XMLName xml.Name `xml:"krt"`
	Value   string   `xml:",chardata"` // (#PCDATA | %inline.html; | %inline.lexml;)*
}
