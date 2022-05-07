package inline_lexml

import "encoding/xml"

type Primary struct {
	InlineLeXML
	XMLName xml.Name `xml:"primary"`
	Value   string   `xml:",chardata"`
}
