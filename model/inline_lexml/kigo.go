package model

import "encoding/xml"

type Kigo struct {
	InlineLeXML
	XMLName xml.Name `xml:"kigo"`
	Value   string   `xml:",chardata"`
}