package model

import "encoding/xml"

type Obliq struct {
	InlineHTML
	XMLName xml.Name `xml:"obliq"`
	Value   string   `xml:",chardata"`
}
