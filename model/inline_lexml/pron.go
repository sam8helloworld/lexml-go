package inline_lexml

import "encoding/xml"

type Pron struct {
	InlineLeXML
	XMLName xml.Name `xml:"pron"`
	Value   string   `xml:",chardata"` // (#PCDATA)
}
