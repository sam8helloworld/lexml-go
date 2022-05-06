package model

import "encoding/xml"

type KKaeri struct {
	InlineHTML
	XMLName xml.Name `xml:"kkaeri"`
	Value   string   `xml:",chardata"`
}
