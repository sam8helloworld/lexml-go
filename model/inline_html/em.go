package inline_html

import "encoding/xml"

type SmallEM struct {
	InlineHTML
	XMLName xml.Name `xml:"em"`
	Type    string   `xml:"type,attr"`
	Class   string   `xml:"class,attr"`
	Value   string   `xml:",chardata"` // (#PCDATA | %inline.html; | %inline.lexml;)*
}

type LargeEM struct {
	InlineHTML
	XMLName xml.Name `xml:"EM"`
	Type    string   `xml:"type,attr"`
	Class   string   `xml:"class,attr"`
	Value   string   `xml:",chardata"` // (#PCDATA | %inline.html; | %inline.lexml;)*
}
