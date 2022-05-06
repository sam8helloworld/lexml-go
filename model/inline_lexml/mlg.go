package model

import "encoding/xml"

type Mlg struct {
	InlineLeXML
	XMLName xml.Name `xml:"mlg"`
	Value   string   `xml:",chardata"`
}
