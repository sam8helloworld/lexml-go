package model

import "encoding/xml"

type Article struct {
	InlineLeXML
	XMLName xml.Name `xml:"article"`
	Value   string   `xml:",chardata"`
}
