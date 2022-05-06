package model

import "encoding/xml"

type Light struct {
	InlineLeXML
	XMLName xml.Name `xml:"light"`
	Value   string   `xml:",chardata"`
}
