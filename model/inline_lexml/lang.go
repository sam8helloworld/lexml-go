package model

import "encoding/xml"

type Lang struct {
	InlineLeXML
	XMLName xml.Name `xml:"lang"`
	Type    string   `xml:"type,attr"`
	Value   string   `xml:",chardata"`
}
