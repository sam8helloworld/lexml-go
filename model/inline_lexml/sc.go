package model

import "encoding/xml"

type Sc struct {
	InlineLeXML
	XMLName xml.Name `xml:"sc"`
	Value   string   `xml:",chardata"`
}
