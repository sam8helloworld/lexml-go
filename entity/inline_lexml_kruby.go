package entity

import "encoding/xml"

type KRuby struct {
	InlineLeXML
	XMLName xml.Name `xml:"kruby"`
	Value   string   `xml:",chardata"` // (#PCDATA | krt | klrt | %inline.html; | %inline.lexml;)*
}
