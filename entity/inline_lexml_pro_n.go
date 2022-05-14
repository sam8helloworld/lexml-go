package entity

import "encoding/xml"

type ProN struct {
	InlineLeXML
	XMLName xml.Name `xml:"pron-n"`
	Value   string   `xml:",chardata"` // (#PCDATA | %inline.html; | %inline.lexml;)*
}
