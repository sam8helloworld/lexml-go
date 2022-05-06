package model

import "encoding/xml"

type Spellout struct {
	InlineHTML
	XMLName xml.Name `xml:"spellout"`
	Type    string   `xml:"type,attr"`
	Org     string   `xml:"org,attr"`
	Value   string   `xml:",chardata"`
}
