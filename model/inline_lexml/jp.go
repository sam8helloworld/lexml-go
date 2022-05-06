package model

import "encoding/xml"

type Jp struct {
	InlineLeXML
	XMLName xml.Name `xml:"jp"`
	Value   string   `xml:",chardata"`
}
