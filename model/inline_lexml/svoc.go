package model

import "encoding/xml"

type Svoc struct {
	InlineLeXML
	XMLName xml.Name `xml:"svoc"`
	Value   string   `xml:",chardata"`
}
