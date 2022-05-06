package model

import "encoding/xml"

type FullForm struct {
	InlineLeXML
	XMLName xml.Name `xml:"fullform"`
	Value   string   `xml:",chardata"`
}
