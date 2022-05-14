package entity

import "encoding/xml"

type Ex struct {
	InlineLeXML
	XMLName xml.Name `xml:"ex"`
	Type    string   `xml:"type,attr"`
	Value   string   `xml:",chardata"` // (#PCDATA | ex-body | ex-trans | ex-src | ex-misc | %inline.html; | %inline.lexml;)*
}
