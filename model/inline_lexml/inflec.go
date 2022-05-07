package inline_lexml

import "encoding/xml"

type Inflec struct {
	InlineLeXML
	XMLName xml.Name `xml:"inflec"`
	Value   string   `xml:",chardata"`
}
