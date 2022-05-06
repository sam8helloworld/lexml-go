package model

import "encoding/xml"

type Inflec struct {
	InlineLeXML
	XMLName xml.Name `xml:"inflec"`
	Value   string   `xml:",chardata"`
}
