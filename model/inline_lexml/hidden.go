package inline_lexml

import "encoding/xml"

type Hidden struct {
	InlineLeXML
	XMLName xml.Name `xml:"hidden"`
	Type    string   `xml:"type,attr"`
	Memo    string   `xml:"memo,attr"`
	Value   string   `xml:",chardata"`
}
