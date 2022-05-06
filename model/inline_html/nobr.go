package model

import "encoding/xml"

type Nobr struct {
	InlineHTML
	XMLName xml.Name `xml:"nobr"`
	Value   string   `xml:",chardata"`
}
