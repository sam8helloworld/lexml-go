package inline_lexml

import "encoding/xml"

type KlOkuri struct {
	InlineLeXML
	XMLName xml.Name `xml:"klokuri"`
	Value   string   `xml:",chardata"` // (#PCDATA | %inline.html; | %inline.lexml;)*
}
