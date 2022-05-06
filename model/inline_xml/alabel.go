package model

import "encoding/xml"

type ALabel struct {
	InlineXML
	XMLName xml.Name `xml:"alabel"`
	Type    string   `xml:"type,attr"`
	Code    string   `xml:"code,attr"`
	Value   string   `xml:",chardata"`
}
