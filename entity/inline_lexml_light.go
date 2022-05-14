package entity

import "encoding/xml"

type Light struct {
	InlineLeXML
	XMLName xml.Name `xml:"light"`
	Value   string   `xml:",chardata"` // (#PCDATA | %inline.html; | %inline.lexml;)*
}
