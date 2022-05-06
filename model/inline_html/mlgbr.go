package model

import "encoding/xml"

type Mlgbr struct {
	InlineHTML
	XMLName xml.Name `xml:"mlgbr"`
	Value   string   `xml:",chardata"`
}
