package model

import "encoding/xml"

type Mlg struct {
	InlineHTML
	XMLName xml.Name `xml:"mlg"`
	Value   string   `xml:",chardata"`
}
