package model

import "encoding/xml"

type ProV struct {
	InlineLeXML
	XMLName xml.Name `xml:"pron-v"`
	Value   string   `xml:",chardata"`
}
