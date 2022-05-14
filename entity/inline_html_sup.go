package entity

import "encoding/xml"

type SmallSup struct {
	InlineHTML
	XMLName xml.Name `xml:"sup"`
	Value   string   `xml:",chardata"` // (#PCDATA | %inline.html; | %inline.lexml;)*
}

type LargeSup struct {
	InlineHTML
	XMLName xml.Name `xml:"SUP"`
	Value   string   `xml:",chardata"` // (#PCDATA | %inline.html; | %inline.lexml;)*
}
