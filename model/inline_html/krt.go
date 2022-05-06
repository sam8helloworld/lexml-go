package model

import "encoding/xml"

type KRt struct {
	InlineHTML
	XMLName xml.Name `xml:"krt"`
	Value   string   `xml:",chardata"`
}
