package inline_lexml

import "encoding/xml"

type Tw struct {
	InlineLeXML
	XMLName xml.Name `xml:"tw"`
	Value   string   `xml:",chardata"` // (#PCDATA | %inline.html; | %inline.lexml;)*
}
