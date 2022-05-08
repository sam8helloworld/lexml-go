package inline_lexml

import "encoding/xml"

type ExSrcAuthor struct {
	InlineLeXML
	XMLName xml.Name `xml:"ex-src-author"`
	Org     string   `xml:"org,attr"`
	Value   string   `xml:",chardata"` // (#PCDATA | %inline.html; | %inline.lexml;)*
}
