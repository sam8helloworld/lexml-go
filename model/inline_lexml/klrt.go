package model

import "encoding/xml"

type KlRt struct {
	InlineLeXML
	XMLName xml.Name `xml:"klrt"`
	Value   string   `xml:",chardata"`
}
