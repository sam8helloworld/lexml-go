package model

import "encoding/xml"

type Memo struct {
	Structure
	XMLName xml.Name `xml:"memo"`
	Type    string   `xml:"type,attr"`
	Value   string   `xml:",chardata"`
}
