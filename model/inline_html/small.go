package inline_html

import "encoding/xml"

type Small struct {
	InlineHTML
	XMLName xml.Name `xml:"small"`
	Value   string   `xml:",chardata"`
}
