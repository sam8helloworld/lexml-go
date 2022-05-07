package inline_lexml

import "encoding/xml"

type Etym struct {
	InlineLeXML
	XMLName xml.Name `xml:"etym"`
	Value   string   `xml:",chardata"`
}
