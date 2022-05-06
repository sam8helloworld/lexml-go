package model

import "encoding/xml"

type ref struct {
	InlineLeXML
	XMLName xml.Name `xml:"ref"`
	Type    string   `xml:"type,attr"`
	RefId   string   `xml:"refid,attr"`
	Id      string   `xml:"id,attr"`
	Value   string   `xml:",chardata"`
}

type Ref struct {
	InlineLeXML
	XMLName xml.Name `xml:"REF"`
	Type    string   `xml:"type,attr"`
	RefId   string   `xml:"refid,attr"`
	Id      string   `xml:"id,attr"`
	Value   string   `xml:",chardata"`
}
