package entity

import "encoding/xml"

type Etym struct {
	InlineLeXML
	XMLName xml.Name `xml:"etym"`
	Value   string   `xml:",chardata"` // (#PCDATA | %inline.html; | %inline.lexml;)*
}
