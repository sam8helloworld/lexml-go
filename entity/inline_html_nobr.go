package entity

import "encoding/xml"

type Nobr struct {
	InlineHTML
	XMLName xml.Name `xml:"nobr"`
}
