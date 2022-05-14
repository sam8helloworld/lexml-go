package entity

import "encoding/xml"

type Primary struct {
	InlineLeXML
	XMLName xml.Name `xml:"primary"`
	Value   string   `xml:",chardata"` // (#PCDATA | %inline.html; | %inline.lexml;)*
}
