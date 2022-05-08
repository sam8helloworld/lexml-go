package structure

import "encoding/xml"

type Th struct {
	Structure
	XMLName xml.Name `xml:"th"`
	ColSpan string   `xml:"colspan,attr"`
	RowSpan string   `xml:"rowspan,attr"`
	Value   string   `xml:",chardata"` //  (#PCDATA | %inline.html; | %inline.lexml;)*
}
