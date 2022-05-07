package inline_lexml

import "encoding/xml"

type Ggk struct {
	InlineLeXML
	XMLName xml.Name `xml:"glabel"`
	Class   string   `xml:"class,attr"`
	Yomi    string   `xml:"yomi,attr"`
	Value   string   `xml:",chardata"`
}
