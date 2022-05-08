package inline_html

import "encoding/xml"

type SmallI struct {
	InlineHTML
	XMLName xml.Name `xml:"i"`
	Type    string   `xml:"type,attr"`
	Class   string   `xml:"class,attr"`
	Value   string   `xml:",chardata"` // (#PCDATA | %inline.html; | %inline.lexml;)*
}

type LargeI struct {
	InlineHTML
	XMLName xml.Name `xml:"I"`
	Type    string   `xml:"type,attr"`
	Class   string   `xml:"class,attr"`
	Value   string   `xml:",chardata"` // (#PCDATA | %inline.html; | %inline.lexml;)*
}
