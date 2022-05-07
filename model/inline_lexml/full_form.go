package inline_lexml

import "encoding/xml"

type FullForm struct {
	InlineLeXML
	XMLName xml.Name `xml:"fullform"`
	Value   string   `xml:",chardata"` // (#PCDATA | %inline.html; | %inline.lexml;)*
}
