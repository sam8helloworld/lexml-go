package entity

import "encoding/xml"

type Inflec struct {
	InlineLeXML
	XMLName xml.Name `xml:"inflec"`
	Value   string   `xml:",chardata"` // (#PCDATA | %inline.html; | %inline.lexml;)*
}
