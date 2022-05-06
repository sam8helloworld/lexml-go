package model

import "encoding/xml"

type KOkuri struct {
	InlineHTML
	XMLName xml.Name `xml:"kokuri"`
	Value   string   `xml:",chardata"`
}
