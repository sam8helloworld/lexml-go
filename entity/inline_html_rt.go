package entity

import "encoding/xml"

type Rt struct {
	InlineHTML
	XMLName xml.Name `xml:"rt"`
	Value   string   `xml:",chardata"` // (#PCDATA | %inline.html; | %inline.lexml;)*
}
