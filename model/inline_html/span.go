package inline_html

import "encoding/xml"

type Span struct {
	InlineHTML
	XMLName xml.Name `xml:"span"`
	Type    string   `xml:"type,attr"`
	Class   string   `xml:"class,attr"`
	Ph      string   `xml:"ph,attr"`
	Value   string   `xml:",chardata"`
}
