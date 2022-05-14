package entity

import "encoding/xml"

type Title struct {
	Structure
	XMLName xml.Name `xml:"title"`
	Type    string   `xml:"type,attr"`
	Value   string   `xml:",chardata"` // (#PCDATA | %inline.html; | %inline.lexml;)*
}
