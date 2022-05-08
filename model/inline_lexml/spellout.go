package inline_lexml

import "encoding/xml"

type Spellout struct {
	InlineLeXML
	XMLName xml.Name `xml:"spellout"`
	Type    string   `xml:"type,attr"`
	Org     string   `xml:"org,attr"`
	Value   string   `xml:",chardata"` // (#PCDATA | %inline.html; | %inline.lexml;)*
}
