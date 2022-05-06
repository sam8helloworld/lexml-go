package model

import "encoding/xml"

type Mlgbr struct {
	InlineLeXML
	XMLName xml.Name `xml:"mlgbr"`
	Value   string   `xml:",chardata"`
}
