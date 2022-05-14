package entity

import "encoding/xml"

type KKaeri struct {
	InlineLeXML
	XMLName xml.Name `xml:"kkaeri"`
	Value   string   `xml:",chardata"` // (#PCDATA | %inline.html; | %inline.lexml;)*
}
