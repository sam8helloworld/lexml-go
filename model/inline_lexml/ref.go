package inline_lexml

import "encoding/xml"

type SmallRef struct {
	InlineLeXML
	XMLName xml.Name `xml:"ref"`
	Type    string   `xml:"type,attr"`
	RefId   string   `xml:"refid,attr"`
	Id      string   `xml:"id,attr"`
	Value   string   `xml:",chardata"` // (#PCDATA | %inline.html; | %inline.lexml;)*
}

type LargeRef struct {
	InlineLeXML
	XMLName xml.Name `xml:"REF"`
	Type    string   `xml:"type,attr"`
	RefId   string   `xml:"refid,attr"`
	Id      string   `xml:"id,attr"`
	Value   string   `xml:",chardata"` // (#PCDATA | %inline.html; | %inline.lexml;)*
}
