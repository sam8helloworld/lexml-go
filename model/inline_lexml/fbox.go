package inline_lexml

import "encoding/xml"

type fbox struct {
	InlineLeXML
	XMLName xml.Name `xml:"fbox"`
	Value   string   `xml:",chardata"`
}

type FBox struct {
	InlineLeXML
	XMLName xml.Name `xml:"FBOX"`
	Value   string   `xml:",chardata"`
}
