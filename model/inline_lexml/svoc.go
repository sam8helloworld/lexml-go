package inline_lexml

import "encoding/xml"

type Svoc struct {
	InlineLeXML
	XMLName xml.Name `xml:"svoc"`
	Value   string   `xml:",chardata"`
}
