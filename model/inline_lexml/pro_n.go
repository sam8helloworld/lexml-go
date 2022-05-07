package inline_lexml

import "encoding/xml"

type ProN struct {
	InlineLeXML
	XMLName xml.Name `xml:"pron-n"`
	Value   string   `xml:",chardata"`
}
