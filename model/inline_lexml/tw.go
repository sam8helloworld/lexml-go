package model

import "encoding/xml"

type Tw struct {
	InlineLeXML
	XMLName xml.Name `xml:"tw"`
	Value   string   `xml:",chardata"`
}
