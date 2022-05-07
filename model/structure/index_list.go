package structure

import "encoding/xml"

type IndexList struct {
	Structure
	XMLName xml.Name `xml:"indexlist"`
	Type    string   `xml:"type,attr"`
	Value   string   `xml:",chardata"` // (#PCDATA | %inline.html; | %inline.lexml;)*
}
