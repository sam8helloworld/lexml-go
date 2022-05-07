package inline_lexml

import "encoding/xml"

type Ex struct {
	InlineLeXML
	XMLName xml.Name `xml:"ex"`
	Type    string   `xml:"type,attr"`
	Value   string   `xml:",chardata"`
}
