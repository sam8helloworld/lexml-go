package model

import "encoding/xml"

type KRt struct {
	InlineLeXML
	XMLName xml.Name `xml:"krt"`
	Value   string   `xml:",chardata"`
}
