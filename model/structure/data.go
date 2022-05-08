package structure

import "encoding/xml"

type Data struct {
	Structure
	XMLName xml.Name `xml:"data"`
	Type    string   `xml:"type,attr"`
	Value   string   `xml:",chardata"` // (#PCDATA | %inline.html; | %inline.lexml;)*
}
