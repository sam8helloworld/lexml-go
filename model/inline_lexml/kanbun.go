package inline_lexml

import "encoding/xml"

type Kanbun struct {
	InlineLeXML
	XMLName xml.Name `xml:"kanbun"`
	Type    string   `xml:"type,attr"`
	SubID   string   `xml:"subid,attr"`
	Value   string   `xml:",chardata"` // (#PCDATA | kkaeri | kruby | kokuri | klokuri | %inline.html; | %inline.lexml;)*
}
