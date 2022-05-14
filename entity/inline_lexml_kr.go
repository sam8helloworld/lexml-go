package entity

import "encoding/xml"

type Kr struct {
	InlineLeXML
	XMLName xml.Name `xml:"kr"`
	Value   string   `xml:",chardata"` // (#PCDATA | %inline.html; | %inline.lexml;)*
}
