package inline_html

import "encoding/xml"

type SmallSub struct {
	InlineHTML
	XMLName xml.Name `xml:"sub"`
	Value   string   `xml:",chardata"`
}

type LargeSub struct {
	InlineHTML
	XMLName xml.Name `xml:"SUB"`
	Value   string   `xml:",chardata"`
}
