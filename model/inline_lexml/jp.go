package inline_lexml

import "encoding/xml"

type Jp struct {
	InlineLeXML
	XMLName xml.Name `xml:"jp"`
	Value   string   `xml:",chardata"` // (#PCDATA | %inline.html; | %inline.lexml;)*
}
