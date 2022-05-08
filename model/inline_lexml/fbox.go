package inline_lexml

import "encoding/xml"

type SmallFBox struct {
	InlineLeXML
	XMLName xml.Name `xml:"fbox"`
	Value   string   `xml:",chardata"` // (#PCDATA | %inline.html; | %inline.lexml;)*
}

type LargeFBox struct {
	InlineLeXML
	XMLName xml.Name `xml:"FBOX"`
	Value   string   `xml:",chardata"` // (#PCDATA | %inline.html; | %inline.lexml;)*
}
