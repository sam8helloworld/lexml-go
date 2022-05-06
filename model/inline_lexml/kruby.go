package model

import "encoding/xml"

type KRuby struct {
	InlineLeXML
	XMLName xml.Name `xml:"kruby"`
	Value   string   `xml:",chardata"`
}
