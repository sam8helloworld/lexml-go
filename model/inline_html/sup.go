package model

import "encoding/xml"

type sup struct {
	InlineHTML
	XMLName xml.Name `xml:"sup"`
	Value   string   `xml:",chardata"`
}

type Sup struct {
	InlineHTML
	XMLName xml.Name `xml:"SUP"`
	Value   string   `xml:",chardata"`
}
