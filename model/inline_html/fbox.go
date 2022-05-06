package model

import "encoding/xml"

type fbox struct {
	InlineHTML
	XMLName xml.Name `xml:"fbox"`
	Value   string   `xml:",chardata"`
}

type FBox struct {
	InlineHTML
	XMLName xml.Name `xml:"FBOX"`
	Value   string   `xml:",chardata"`
}
