package inline_html

import "encoding/xml"

type Big struct {
	InlineHTML
	XMLName xml.Name `xml:"big"`
	Type    string   `xml:"type,attr"`
	Class   string   `xml:"class,attr"`
	Value   string   `xml:",chardata"` // (#PCDATA | %inline.html; | %inline.lexml;)*
}
