package model

import "encoding/xml"

type KRuby struct {
	InlineHTML
	XMLName xml.Name `xml:"kruby"`
	Value   string   `xml:",chardata"`
}
