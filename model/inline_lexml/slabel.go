package inline_lexml

import "encoding/xml"

type SLabel struct {
	InlineLeXML
	XMLName xml.Name `xml:"slabel"`
	Type    string   `xml:"type,attr"`
	Genre   string   `xml:"genre,attr"`
	Code    string   `xml:"code,attr"`
	Value   string   `xml:",chardata"`
}
