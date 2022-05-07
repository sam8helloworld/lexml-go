package inline_lexml

import "encoding/xml"

type Kr struct {
	InlineLeXML
	XMLName xml.Name `xml:"kr"`
	Value   string   `xml:",chardata"`
}
