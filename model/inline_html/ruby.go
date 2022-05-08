package inline_html

import "encoding/xml"

type Ruby struct {
	InlineHTML
	XMLName xml.Name `xml:"ruby"`
	Rb
	Rt
}
