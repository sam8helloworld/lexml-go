package inline_html

import "encoding/xml"

type SmallSup struct {
	InlineHTML
	XMLName xml.Name `xml:"sup"`
	Value   string   `xml:",chardata"`
}

type LargeSup struct {
	InlineHTML
	XMLName xml.Name `xml:"SUP"`
	Value   string   `xml:",chardata"`
}
