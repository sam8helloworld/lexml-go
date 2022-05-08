package inline_lexml

import "encoding/xml"

type IPA struct {
	InlineLeXML
	XMLName xml.Name `xml:"ipa"`
	Value   string   `xml:",chardata"` // (#PCDATA | %inline.html; | %inline.lexml;)*
}
