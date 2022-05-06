package model

import "encoding/xml"

type Abbr struct {
	InlineLeXML
	XMLName xml.Name `xml:"abbr"`
	Value   string   `xml:",chardata"`
}
