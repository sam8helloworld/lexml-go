package model

import "encoding/xml"

type KlRt struct {
	InlineHTML
	XMLName xml.Name `xml:"klrt"`
	Value   string   `xml:",chardata"`
}
