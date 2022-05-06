package model

import "encoding/xml"

type Sc struct {
	InlineHTML
	XMLName xml.Name `xml:"sc"`
	Value   string   `xml:",chardata"`
}
