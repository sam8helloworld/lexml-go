package inline_lexml

import "encoding/xml"

type ExSrcTitle struct {
	InlineLeXML
	XMLName xml.Name `xml:"ex-src-title"`
	Org     string   `xml:"org,attr"`
	Value   string   `xml:",chardata"` // (#PCDATA | %inline.html; | %inline.lexml;)*
}
