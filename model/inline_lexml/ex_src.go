package model

import "encoding/xml"

type ExSrc struct {
	InlineLeXML
	XMLName xml.Name `xml:"ex-src"`
	RefID   string   `xml:"refid,attr"`
	Type    string   `xml:"type,attr"`
	Value   string   `xml:",chardata"`
}
