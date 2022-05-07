package inline_lexml

import "encoding/xml"

type ExMisc struct {
	InlineLeXML
	XMLName xml.Name `xml:"ex-misc"`
	Type    string   `xml:"type,attr"`
	Value   string   `xml:",chardata"`
}
