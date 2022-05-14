package entity

import "encoding/xml"

type Pinyin struct {
	InlineLeXML
	XMLName xml.Name `xml:"pinyin"`
	Value   string   `xml:",chardata"` // (#PCDATA | %inline.html; | %inline.lexml;)*
}
