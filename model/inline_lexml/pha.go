package inline_lexml

import "encoding/xml"

type PhA struct {
	InlineLeXML
	XMLName xml.Name `xml:"pha"`
	Type    string   `xml:"type,attr"`
	Value   string   `xml:",chardata"`
}
