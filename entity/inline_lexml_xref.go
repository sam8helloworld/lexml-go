package entity

import "encoding/xml"

type SmallXref struct {
	InlineLeXML
	XMLName xml.Name `xml:"xref"`
	Type    string   `xml:"type,attr"`
	RefId   string   `xml:"refid,attr"`
	Id      string   `xml:"id,attr"`
	Value   string   `xml:",chardata"` // (#PCDATA | %inline.html; | %inline.lexml;)*
}

type LargeXRef struct {
	InlineLeXML
	XMLName xml.Name `xml:"XREF"`
	Type    string   `xml:"type,attr"`
	RefId   string   `xml:"refid,attr"`
	Id      string   `xml:"id,attr"`
	Value   string   `xml:",chardata"` // (#PCDATA | %inline.html; | %inline.lexml;)*
}
