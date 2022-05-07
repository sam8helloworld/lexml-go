package inline_lexml

import "encoding/xml"

type SmallPos struct {
	InlineLeXML
	XMLName xml.Name `xml:"pos"`
	Value   string   `xml:",chardata"` // (#PCDATA)
}

type LargePos struct {
	InlineLeXML
	XMLName xml.Name `xml:"Pos"`
	Value   string   `xml:",chardata"` // (#PCDATA)
}
