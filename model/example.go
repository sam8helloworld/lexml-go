package model

import "encoding/xml"

type Example struct {
	XMLName   xml.Name `xml:"example"`
	SubID     string   `xml:"subid,attr"`
	Type      string   `xml:"type,attr"`
	Delimiter string   `xml:"delimiter,attr"`
	Value     string   `xml:",innerxml"`
	// TODO: ex-body, ex-trans, ex-src, ex-miscを追加
}
