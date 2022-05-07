package inline_lexml

import "encoding/xml"

type KOkuri struct {
	InlineLeXML
	XMLName xml.Name `xml:"kokuri"`
	Value   string   `xml:",chardata"` // (#PCDATA | %inline.html; | %inline.lexml;)*
}
