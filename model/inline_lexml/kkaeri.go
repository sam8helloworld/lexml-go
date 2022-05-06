package model

import "encoding/xml"

type KKaeri struct {
	InlineLeXML
	XMLName xml.Name `xml:"kkaeri"`
	Value   string   `xml:",chardata"`
}
