package inline_lexml

import "encoding/xml"

type Cn struct {
	InlineLeXML
	XMLName xml.Name `xml:"cn"`
	Value   string   `xml:",chardata"`
}
