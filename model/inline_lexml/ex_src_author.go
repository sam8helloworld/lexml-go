package model

import "encoding/xml"

type ExSrcAuthor struct {
	InlineLeXML
	XMLName xml.Name `xml:"ex-src-author"`
	Org     string   `xml:"org,attr"`
	Value   string   `xml:",chardata"`
}
