package model

import "encoding/xml"

type ExSrcTitle struct {
	InlineLeXML
	XMLName xml.Name `xml:"ex-src-title"`
	Org     string   `xml:"org,attr"`
	Value   string   `xml:",chardata"`
}
