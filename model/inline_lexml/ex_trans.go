package model

import "encoding/xml"

type ExTrans struct {
	InlineLeXML
	XMLName xml.Name `xml:"ex-trans"`
	Lang    string   `xml:"lang,attr"`
	Type    string   `xml:"type,attr"`
	Value   string   `xml:",chardata"`
}
