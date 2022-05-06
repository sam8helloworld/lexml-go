package model

import "encoding/xml"

type Rb struct {
	InlineHTML
	XMLName xml.Name `xml:"rb"`
	Value   string   `xml:",chardata"`
}
