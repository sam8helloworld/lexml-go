package model

import "encoding/xml"

type KlOkuri struct {
	InlineLeXML
	XMLName xml.Name `xml:"klokuri"`
	Value   string   `xml:",chardata"`
}
