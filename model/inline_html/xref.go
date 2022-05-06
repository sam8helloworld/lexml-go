package model

import "encoding/xml"

type xref struct {
	InlineHTML
	XMLName xml.Name `xml:"xref"`
	Type    string   `xml:"type,attr"`
	RefId   string   `xml:"refid,attr"`
	Id      string   `xml:"id,attr"`
	Value   string   `xml:",chardata"`
}

type XRef struct {
	InlineHTML
	XMLName xml.Name `xml:"XREF"`
	Type    string   `xml:"type,attr"`
	RefId   string   `xml:"refid,attr"`
	Id      string   `xml:"id,attr"`
	Value   string   `xml:",chardata"`
}