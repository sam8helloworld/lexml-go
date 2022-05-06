package model

import "encoding/xml"

type IPA struct {
	InlineLeXML
	XMLName xml.Name `xml:"ipa"`
	Value   string   `xml:",chardata"`
}
