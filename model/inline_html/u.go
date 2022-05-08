package inline_html

import "encoding/xml"

type SmallU struct {
	InlineHTML
	XMLName xml.Name `xml:"u"`
	Type    string   `xml:"type,attr"`
	Class   string   `xml:"class,attr"`
	Value   string   `xml:",chardata"` // (#PCDATA | %inline.html; | %inline.lexml;)*
}

type LargeU struct {
	InlineHTML
	XMLName xml.Name `xml:"U"`
	Type    string   `xml:"type,attr"`
	Class   string   `xml:"class,attr"`
	Value   string   `xml:",chardata"` // (#PCDATA | %inline.html; | %inline.lexml;)*
}
