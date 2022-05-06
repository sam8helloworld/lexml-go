package model

import "encoding/xml"

type Vnum struct {
	InlineLeXML
	XMLName xml.Name `xml:"vnum"`
	Type    string   `xml:"type,attr"`
	Value   string   `xml:",chardata"`
}
