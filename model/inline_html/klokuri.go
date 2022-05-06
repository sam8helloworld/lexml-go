package model

import "encoding/xml"

type KlOkuri struct {
	InlineHTML
	XMLName xml.Name `xml:"klokuri"`
	Value   string   `xml:",chardata"`
}
