package model

import "encoding/xml"

type gender struct {
	InlineLeXML
	XMLName xml.Name `xml:"gender"`
	Value   string   `xml:",chardata"`
}

type Gender struct {
	InlineLeXML
	XMLName xml.Name `xml:"Gender"`
	Value   string   `xml:",chardata"`
}
