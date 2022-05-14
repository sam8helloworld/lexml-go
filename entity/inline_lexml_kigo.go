package entity

import "encoding/xml"

type Kigo struct {
	InlineLeXML
	XMLName xml.Name `xml:"kigo"`
	Value   string   `xml:",chardata"` // (#PCDATA | %inline.html; | %inline.lexml;)*
}
