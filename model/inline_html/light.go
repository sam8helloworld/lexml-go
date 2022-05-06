package model

import "encoding/xml"

type Light struct {
	InlineHTML
	XMLName xml.Name `xml:"light"`
	Value   string   `xml:",chardata"`
}
