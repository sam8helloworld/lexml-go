package structure

import (
	"encoding/xml"
)

type Headword struct {
	Structure
	XMLName   xml.Name `xml:"headword"`
	Type      string   `xml:"type,attr"`
	Delimiter string   `xml:"delimiter,attr"`
	Value     string   `xml:",chardata"` // (#PCDATA | %inline.html; | %inline.lexml;)*
}
