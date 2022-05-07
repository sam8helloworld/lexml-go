package inline_lexml

import "encoding/xml"

type SmallGender struct {
	InlineLeXML
	XMLName xml.Name `xml:"gender"`
	Value   string   `xml:",chardata"` // (#PCDATA)
}

type LargeGender struct {
	InlineLeXML
	XMLName xml.Name `xml:"Gender"`
	Value   string   `xml:",chardata"` // (#PCDATA)
}
