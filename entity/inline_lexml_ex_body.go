package entity

import "encoding/xml"

type ExBody struct {
	InlineLeXML
	XMLName xml.Name `xml:"ex-body"`
	Lang    string   `xml:"lang,attr"`
	Type    string   `xml:"type,attr"`
	Value   string   `xml:",chardata"` // (#PCDATA | %inline.html; | %inline.lexml;)*
}
