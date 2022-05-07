package inline_lexml

import "encoding/xml"

type Pinyin struct {
	InlineLeXML
	XMLName xml.Name `xml:"pinyin"`
	Value   string   `xml:",chardata"`
}
