package model

import "encoding/xml"

type sub struct {
	InlineHTML
	XMLName xml.Name `xml:"sub"`
	Value   string   `xml:",chardata"`
}

type Sub struct {
	InlineHTML
	XMLName xml.Name `xml:"SUB"`
	Value   string   `xml:",chardata"`
}
