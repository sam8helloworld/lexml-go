package inline_lexml

import "encoding/xml"

type Gix struct {
	InlineLeXML
	XMLName xml.Name `xml:"gix"`
	Alt     string   `xml:"alt,attr"`
	Type    string   `xml:"type,attr"`
	AltImg  string   `xml:"altimg,attr"`
	Value   string   `xml:",chardata"`
}
